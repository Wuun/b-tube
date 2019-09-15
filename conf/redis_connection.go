package conf

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisConnection is the singleton of redis connection.
var RedisConnection *redis.Client

// Redis init the RedisConnection.
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	RedisConnection = client
}
