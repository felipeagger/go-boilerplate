package cache

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	onceCache sync.Once
	cacheClientSvc Service
)

func InitCacheClientSvc(cacheHost string, cachePort string, cachePassword string) {
	onceCache.Do(func() {

		cacheSvc, err := NewCacheService(cacheHost, cachePort, cachePassword)

		if err != nil {
			panic(err)
		}

		cacheClientSvc = cacheSvc
	})
}

func GetCacheClient() Service {
	return cacheClientSvc
}


type Client struct {
	Client    *redis.Client
}

func NewSimpleCacheClient(host, port, password string) (*Client, error) {

	client, err := NewRedisClient(host, port, password)
	if err != nil {
		return nil, err
	}

	cache := &Client{
		Client: client,
	}

	return cache, nil
}

func (cache *Client) HSet(ctx context.Context, key string, values ...interface{}) error {
	return cache.Client.HSet(ctx, key, values).Err()
}

func (cache *Client) Set(ctx context.Context, key string, values interface{}, expiration time.Duration) error {
	return cache.Client.Set(ctx, key, values, expiration).Err()
}

func (cache *Client) HSetExp(ctx context.Context, key string, expiration time.Duration, values ...interface{}) error {
	pipe := cache.Client.Pipeline()
	pipe.HSet(ctx, key, values)
	pipe.Expire(ctx, key, expiration)
	_, err := pipe.Exec(ctx)
	return err
}

func (cache *Client) HSetNX(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) (set bool, err error) {

	set, err = cache.Client.HSetNX(ctx, key, field, value).Result()
	if set == true {
		err = cache.Expire(ctx, key, expiration)
	}

	return set, err
}

func (cache *Client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return cache.Client.Expire(ctx, key, expiration).Err()
}

func (cache *Client) Get(ctx context.Context, key string) (string, error) {

	data, err := cache.Client.Get(ctx, key).Result()
	if err != nil && err.Error() == "redis: nil" {
		return data, nil
	}

	return data, err
}

func (cache *Client) HGet(ctx context.Context, key string, field string) (string, error) {

	data, err := cache.Client.HGet(ctx, key, field).Result()
	if err != nil && err.Error() == "redis: nil" {
		return data, nil
	}

	return data, err
}

func (cache *Client) HGetAll(ctx context.Context, key string) map[string]string {
	return cache.Client.HGetAll(ctx, key).Val()
}

func (cache *Client) HDel(ctx context.Context, key string, fields string) error {
	return cache.Client.HDel(ctx, key, fields).Err()
}

func (cache *Client) Del(ctx context.Context, key string) error {
	return cache.Client.Del(ctx, key).Err()
}

func (cache *Client) Pipeline() redis.Pipeliner {
	return cache.Client.Pipeline()
}

// NewCacheClient return a new instance of cache client
func NewRedisClient(hostname, port, password string) (*redis.Client, error) {

	var client *redis.Client

	cachePort := "6379"
	if port != "" {
		cachePort = port
	}

	if len(password) > 0 {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", hostname, cachePort),
			DB:       0, // use default DB
			Password: password,
			TLSConfig: &tls.Config{
				RootCAs: nil,
			},
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", hostname, cachePort),
			DB:   0, // use default DB
		})
	}

	client.AddHook(redisotel.TracingHook{})

	_, err := client.Ping(context.TODO()).Result()

	if err != nil {
		println("ERROR ON REDIS: NewCacheClient()")
		return nil, err
	}

	return client, nil
}
