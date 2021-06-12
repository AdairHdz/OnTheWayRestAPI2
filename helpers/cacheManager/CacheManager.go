package cacheManager

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:8250",
		Password: "",
		DB:       0,
	})
}

func Save(key string, value interface{}, expiration time.Duration) error {
	redisStatus := rdb.Set(context.Background(), key, value, expiration)
	_, err := redisStatus.Result()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) ([]byte, error) {
	redisStatus := rdb.Get(context.Background(), key)
	result, err := redisStatus.Result()
	resultParsedInBytes := []byte(result)

	return resultParsedInBytes, err
}
