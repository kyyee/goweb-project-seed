package xdatasource

import (
	"fmt"
	"goweb-project-seed/framework/core/xerrors"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(c *Config) (*gorm.DB, error) {
	// todo 密码使用enc加密
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.Port, c.User, c.Password, c.DbName)
	db, err := gorm.Open(postgres.Open(dsn))
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
		return db.Debug(), nil
	}
	return db, nil
}

func NewMariadb(c *Config) (*gorm.DB, error) {
	// todo 密码使用enc加密
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
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
		return db.Debug(), nil
	}
	return db, nil
}

func NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
