package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Service interface {
	HSet(ctx context.Context, key string, values ...interface{}) error
	Set(ctx context.Context, key string, values interface{}, expiration time.Duration) error
	HSetExp(ctx context.Context, key string, expiration time.Duration, values ...interface{}) error
	HSetNX(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) (set bool, err error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) map[string]string
	HDel(ctx context.Context, key string, fields string) error
	Del(ctx context.Context, key string) error
	Pipeline() redis.Pipeliner
}

func NewCacheService(host, port, password string) (Service, error) {
	return NewSimpleCacheClient(host, port, password)
}
