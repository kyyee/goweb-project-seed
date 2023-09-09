package xgorm

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDataSource(t *testing.T) {
	ctx := context.Background()
	cfg := DefaultConfig()
	cfg.DSN = "host=localhost user=postgres password=123456 dbname=gps port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	cfg.Debug = true
	db, err := New(cfg)
	assert.Nil(t, err)
	s, err := db.DB()
	assert.Nil(t, err)
	err = s.Ping()
	assert.Nil(t, err)
	var user User
	db.WithContext(ctx).Model(&User{}).Where("id=1").Find(&user)

	var user2 User
	db.WithContext(ctx).Model(&User{}).Where("name=?", "test").Find(&user2)
}

type User struct {
	Id   int64  `grom:"column:id"`
	Name string `grom:"column:name"`
}

func (user *User) TableName() string {
	return "user"
}

func TestDecryptPassword(t *testing.T) {
	dsn := "host=localhost user=postgres password=ENC(123456) dbname=gps port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	log.Printf(dsn)
}
