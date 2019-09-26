package conf

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

//RedisConnect is sigleton of redis connection.
var RedisConnect *redis.Client

//Redis init redis connection.
func Redis() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"), // no password set
		DB:       db,                    // use default DB
	})

	RedisConnect = client
}
