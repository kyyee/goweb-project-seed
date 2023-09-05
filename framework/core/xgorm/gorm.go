package xgorm

import (
	"goweb-project-seed/framework/core/xerrors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	s.SetMaxIdleConns(c.MaxIdleConns)
	s.SetMaxOpenConns(c.MaxOpenConns)
	if c.Debug {
		db = db.Debug()
	}
	if len(c.SlaveDSN) > 0 {
		list := make([]gorm.Dialector, 0)
		for _, itemDSN := range c.SlaveDSN {
			list = append(list, postgres.Open(itemDSN))
		}
	}

	return db, nil

}
