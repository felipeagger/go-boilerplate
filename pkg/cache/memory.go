package cache

import (
	"context"
	"errors"
	"github.com/dgraph-io/ristretto"
	"sync"
	"time"
)

var (
	onceCacheMem sync.Once
	cacheMemSvc *ClientMem
)

func InitCacheMemSvc() {
	onceCache.Do(func() {

		cacheSvc, err := NewCacheMemClient()

		if err != nil {
			panic(err)
		}

		cacheMemSvc = cacheSvc
	})
}

func GetCacheMemClient() Service {
	return cacheMemSvc
}


type ClientMem struct {
	Client    *ristretto.Cache
}

func NewCacheMemClient() (*ClientMem, error) {

	client, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	cache := &ClientMem{
		Client: client,
	}

	return cache, err
}


func (cache *ClientMem) Set(ctx context.Context, key string, values interface{}, expiration time.Duration) error {
	ok := cache.Client.SetWithTTL(key, values, 0, expiration)
	cache.Client.Wait()

	if !ok {
		return errors.New("error on set")
	}
	return nil
}

func (cache *ClientMem) Get(ctx context.Context, key string) (interface{}, error) {

	data, ok := cache.Client.Get(key)
	if !ok {
		return nil, errors.New("error on get")
	}

	return data, nil
}

func (cache *ClientMem) Del(ctx context.Context, key string) error {
	cache.Client.Del(key)
	return nil
}

func (cache *ClientMem) HSet(ctx context.Context, key string, values ...interface{}) error {
	return nil
}

func (cache *ClientMem) HSetExp(ctx context.Context, key string, expiration time.Duration, values ...interface{}) error {
	return nil
}

func (cache *ClientMem) HSetNX(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) (set bool, err error) {
	return false, nil
}

func (cache *ClientMem) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return nil
}

func (cache *ClientMem) HGet(ctx context.Context, key string, field string) (string, error) {
	return "", nil
}

func (cache *ClientMem) HGetAll(ctx context.Context, key string) map[string]string {
	return map[string]string{}
}

func (cache *ClientMem) HDel(ctx context.Context, key string, fields string) error {
	return nil
}