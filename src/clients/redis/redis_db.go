package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

const(
	redisOauthHost     = "DB_HOST"
	redisOauthPassword = "DB_PASSWORD"
	redisOauthDb       = "DB_SCHEMA"
)

var (
	rdb *redis.Client
	addr     = os.Getenv(redisOauthHost)
        password = os.Getenv(redisOauthPassword)
        db,_     = strconv.Atoi(os.Getenv(redisOauthDb))
)

func init() {
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
