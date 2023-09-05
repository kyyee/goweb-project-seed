package xdatasource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgres(t *testing.T) {
	db, err := NewPostgres(&Config{
		Debug:        true,
		User:         "postgres",
		Password:     "123456",
		Host:         "localhost",
		Port:         5432,
		DbName:       "gps",
		MaxIdleConns: 10,
		MaxOpenConns: 5,
	})
	assert.Nil(t, err)
	s, err := db.DB()
	assert.Nil(t, err)
	err = s.Ping()
	assert.Nil(t, err)
}

func TestNewMariadb(t *testing.T) {
	db, err := NewMariadb(&Config{
		Debug:        true,
		User:         "root",
		Password:     "123456",
		Host:         "localhost:3306",
		DbName:       "gps",
		MaxIdleConns: 10,
		MaxOpenConns: 5,
	})
	assert.Nil(t, err)
	s, err := db.DB()
	assert.Nil(t, err)
	err = s.Ping()
	assert.Nil(t, err)
}
