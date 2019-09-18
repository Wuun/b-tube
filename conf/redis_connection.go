package conf

import (
	"strconv"

	"github.com/go-redis/redis"
)

// RedisConnection is the singleton of redis connection.
var RedisConnection *redis.Client

// Redis init the RedisConnection.
func Redis() {
	db, _ := strconv.ParseUint(GlobalConf.RedisDB, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     GlobalConf.RedisAddr,
		Password: GlobalConf.RedisPW,
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisConnection = client
}
