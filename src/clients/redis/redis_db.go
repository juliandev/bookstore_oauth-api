package redis

import (
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetSession() (*redis.Client, error) {
	if err := rdb.Ping(rdb.Context()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
