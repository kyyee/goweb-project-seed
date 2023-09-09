package xgorm

import (
	"database/sql/driver"
	"fmt"
	"goweb-project-seed/framework/core/utils"
	"goweb-project-seed/framework/core/xerrors"
	"log"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Config struct {
	Debug           bool          `json:"debug" mapstructure:"debug"`
	DSN             string        `json:"dsn" mapstructure:"dsn"`
	SlaveDSN        []string      `json:"slave_dsn" mapstructure:"slave_dsn"`
	MaxIdleConns    int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns" mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" mapstrucure:"conn_max_lifetime"`
	hooks           []Hook
}

func DefaultConfig() *Config {
	return &Config{
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: utils.Duration("300s"),
	}
}

type Handler func(db *gorm.DB)
type Processor interface {
	Get(name string) func(*gorm.DB)
	Replace(name string, handler func(*gorm.DB)) error
}

type Hook func(*Config) func(next Handler) Handler

func New(c *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DSN))
	if err != nil {
		return nil, xerrors.NewAssignCodeError(xerrors.CONNECTION_FAILED)
	}
	s, err := db.DB()
	if err != nil {
		return nil, xerrors.NewAssignCodeError(xerrors.DATA_SOURCE_ERROR)
	}
	if len(c.SlaveDSN) > 0 {
		list := make([]gorm.Dialector, 0)
		for _, itemDSN := range c.SlaveDSN {
			list = append(list, postgres.Open(itemDSN))
		}
		resolver := dbresolver.Register(dbresolver.Config{
			Replicas: list,
		})
		resolver.SetMaxIdleConns(c.MaxIdleConns)
		resolver.SetMaxOpenConns(c.MaxOpenConns)
		if c.ConnMaxLifetime != 0 {
			resolver.SetConnMaxLifetime(c.ConnMaxLifetime)
		}
		if err := db.Use(resolver); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	s.SetMaxIdleConns(c.MaxIdleConns)
	s.SetMaxOpenConns(c.MaxOpenConns)
	if c.ConnMaxLifetime != 0 {
		s.SetConnMaxLifetime(c.ConnMaxLifetime)
	}
	if c.Debug {
		db = db.Debug()
		c.hooks = append(c.hooks, debugHook)
	}
	replace := func(processor Processor, callback string, hooks ...Hook) {
		handler := processor.Get(callback)
		for _, h := range c.hooks {
			handler = h(c)(handler)
		}
		processor.Replace(callback, handler)
	}

	replace(db.Callback().Create(), "gorm:create", c.hooks...)
	replace(db.Callback().Update(), "gorm:update", c.hooks...)
	replace(db.Callback().Delete(), "gorm:delete", c.hooks...)
	replace(db.Callback().Query(), "gorm:query", c.hooks...)
	replace(db.Callback().Raw(), "gorm:raw", c.hooks...)
	replace(db.Callback().Row(), "gorm:row", c.hooks...)

	return db, nil
}

func debugHook(c *Config) func(Handler) Handler {
	return func(next Handler) Handler {
		return func(db *gorm.DB) {
			begin := time.Now()
			next(db)
			cost := time.Since(begin)
			sqlString := printSql(db.Statement.SQL.String(), db.Statement.Vars, true)
			if db.Error != nil {
				log.Printf(fmt.Sprintf("%s %s %s %s => %s",
					utils.Green("[xgorm response]"),
					utils.Blue(c.DSN),
					utils.Red(utils.TimeToMicroSecondStr(cost)),
					utils.Yellow(fmt.Sprintf("%v", sqlString)),
					utils.Green(fmt.Sprintf("%v", db.Error.Error())),
				))

			} else {
				log.Printf(fmt.Sprintf("%s %s %s %s => %s",
					utils.Green("[xgorm response]"),
					utils.Blue(c.DSN),
					utils.Red(utils.TimeToMicroSecondStr(cost)),
					utils.Yellow(fmt.Sprintf("%v", sqlString)),
					utils.Green(fmt.Sprintf("%v", db.Statement.Dest)),
				))
			}
		}
	}
}

func printSql(sql string, args []interface{}, withArgs bool) string {
	if withArgs {
		return bindSql(sql, args)
	}
	return sql
}

func bindSql(originSql string, args []interface{}) (sql string) {
	formattedValues := make([]string, 0)
	for _, value := range args {
		indirectValue := reflect.Indirect(reflect.ValueOf(value))
		if indirectValue.IsValid() {
			value = indirectValue.Interface()
			if t, ok := value.(time.Time); ok {
				formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2000-01-01 12:00:00")))
			} else if b, ok := value.([]byte); ok {
				if str := string(b); isPrint(str) {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
				} else {
					formattedValues = append(formattedValues, "'<binary>'")
				}
			} else if r, ok := value.(driver.Valuer); ok {
				if value, err := r.Value(); err == nil && value != nil {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			} else {
				switch value.(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
					formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
				default:
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			}
		} else {
			formattedValues = append(formattedValues, "NULL")
		}
	}
	if regexp.MustCompile(`\$\d+`).MatchString(originSql) {
		for i, v := range formattedValues {
			placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, i+1)
			sql = regexp.MustCompile(placeholder).ReplaceAllString(originSql, v+"$1")

		}
	} else {
		formattedValuesLen := len(formattedValues)
		for i, v := range regexp.MustCompile(`\?`).Split(originSql, -1) {
			sql += v
			if i < formattedValuesLen {
				sql += formattedValues[i]
			}
		}
	}
	return
}

func isPrint(str string) bool {
	for _, s := range str {
		if !unicode.IsPrint(s) {
			return false
		}
	}
	return true
}
