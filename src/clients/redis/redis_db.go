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

	if err := rdb.Ping(rdb.Context()).Err(); err != nil {
		panic(err)
        }
}

func GetSession() *redis.Client {
	return rdb
}
