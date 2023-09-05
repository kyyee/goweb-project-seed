package xredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (c *Client) Client() redis.Cmdable {
	return c.client
}

func (c *Client) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return c.client.Set(ctx, key, value, expire).Err()
}

func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	i, e := c.client.Exists(ctx, key).Result()
	if e != nil {
		return i == 1, e
	}
	return i == 1, nil
}

func (c *Client) SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	return c.client.SetNX(ctx, key, value, expire).Result()
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Client) Delete(ctx context.Context, key string) (bool, error) {
	i, e := c.client.Exists(ctx, key).Result()
	if e != nil {
		return i == 1, e
	}
	return i == 1, nil
}
