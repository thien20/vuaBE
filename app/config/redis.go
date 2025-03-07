package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(conStr string) *redis.Client {

	// This is for local machine
	// "redis": "localhost:6379" or	redis://<user>:<pass>@localhost:6379/<db>
	// This is for docker
	// redis://<user>:<pass>@reids:6379/<db>

	opt, err := redis.ParseURL(conStr)
	if err != nil {
		panic(err)
	}
	log.Println("Connecting to Redis:", conStr)
	client := redis.NewClient(opt)

	return client
}
