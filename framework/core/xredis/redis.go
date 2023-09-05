package xredis

import (
	"context"
	"errors"
	"goweb-project-seed/framework/core/utils"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addrs        []string      `json:"addrs" mapstructure:"addrs"`
	Password     string        `json:"password" mapstructure:"password"`
	PoolSize     int           `json:"poolSize" mapstructure:"poolSize"`
	MaxRetries   int           `json:"max_retries" mapstructure:"max_retries"`
	MaxRedirects int           `json:"max_redirects" mapstructure:"max_redirects"`
	MinIdleConns int           `json:"min_idle_conns" mapstructure:"min_idle_conns"`
	DialTimeout  time.Duration `json:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"`
}

func DefaultConfig() *Config {
	return &Config{
		PoolSize:     20,
		MaxRetries:   3,
		MaxRedirects: 3,
		MinIdleConns: 100,
		DialTimeout:  utils.Duration("1s"),
		ReadTimeout:  utils.Duration("1s"),
		WriteTimeout: utils.Duration("1s"),
		IdleTimeout:  utils.Duration("120s"),
	}
}

type Client struct {
	client redis.Cmdable
}

func New(c *Config) (*Client, error) {
	if len(c.Addrs) < 1 {
		return nil, errors.New("redis node should not empty")
	} else if len(c.Addrs) == 1 {
		client := redis.NewClient(&redis.Options{
			Addr:         c.Addrs[0],
			PoolSize:     c.PoolSize,
			Password:     c.Password, // todo 密码使用enc加密
			MaxRetries:   c.MaxRetries,
			MinIdleConns: c.MinIdleConns,
			DialTimeout:  c.DialTimeout,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			IdleTimeout:  c.IdleTimeout,
		})
		if e := client.Ping(context.Background()).Err(); e != nil {
			log.Fatalf("init redis cluster client error:%s", e.Error())
		}
		return &Client{client: client}, nil
	} else {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        c.Addrs,
			PoolSize:     c.PoolSize,
			Password:     c.Password,
			MaxRetries:   c.MaxRetries,
			MaxRedirects: c.MaxRedirects,
			MinIdleConns: c.MinIdleConns,
			DialTimeout:  c.DialTimeout,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			IdleTimeout:  c.IdleTimeout,
		})
		if e := client.Ping(context.Background()).Err(); e != nil {
			log.Fatalf("init redis cluster client error:%s", e.Error())
		}
		return &Client{client: client}, nil
	}

}
