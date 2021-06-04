package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"fmt"
)

var (
	rdb *redis.Client
)

func init() {
	addr     := os.Getenv("redis_oauth_host")
	password := os.Getenv("redis_oauth_password")
	db,_     := strconv.Atoi(os.Getenv("redis_oauth_db"))

	fmt.Println(addr, password, db)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := rdb.Ping(rdb.Context()).Err(); err != nil {
		panic(err)
        }
}

func GetSession() *redis.Client {
	return rdb
}
