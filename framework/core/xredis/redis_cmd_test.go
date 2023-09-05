package xredis

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testRedisClient *Client

func InitClient() {
	if testRedisClient != nil {
		return
	}
	cfg := DefaultConfig()
	cfg.Addrs = []string{"127.0.0.1:6379"}
	cfg.Password = "redis_123"
	var err error
	testRedisClient, err = New(cfg)
	if err != nil {
		log.Fatal("redis client init faild:", err.Error())
	}
}

func TestNewClient(t *testing.T) {
	InitClient()
	assert.NotNil(t, testRedisClient)
}

func TestCmds(t *testing.T) {
	InitClient()

	ctx := context.Background()
	err := testRedisClient.Set(ctx, "local:test:key:set", "123456", 5*time.Second)
	assert.Nil(t, err)

	exists, err := testRedisClient.Exists(ctx, "local:test:key:set")
	assert.Nil(t, err)
	assert.True(t, exists)

	result, err := testRedisClient.SetNX(ctx, "local:test:key:set", "123456", 5*time.Second)
	assert.Nil(t, err)
	assert.False(t, result)

	value, err := testRedisClient.Get(ctx, "local:test:key:set")
	assert.Nil(t, err)
	assert.Equal(t, value, "123456")

	result, err = testRedisClient.Delete(ctx, "local:test:key:set")
	assert.Nil(t, err)
	assert.True(t, result)
}
