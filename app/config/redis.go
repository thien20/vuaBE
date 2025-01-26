package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// type client struct {
// 	redisClient *redis.Client
// }

func NewRedis() (*redis.Client, error) {
	ctx := context.Background()
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // use default Addr
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // use default Addr
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })
	// // another way to connect to Redis
	// opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	// if err != nil {
	// 	panic(err)
	// }

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis: ", err.Error())
	}
	return client, nil
}

// // Get title / id from news with Redis
// func (c *client) GetCate(category string) error {
// 	ctx := context.Background()

// 	category, err := c.redisClient.Get(ctx, category).Result()
// 	if err != nil {
// 		return err
// 	}
// 	log.Println("Category: ", category)
// 	return nil
// }

// // Set title / id from news with Redis
// func (c *client) SetCate(category string) error {
// 	ctx := context.Background()

// 	err := c.redisClient.Set(ctx, "the_gioi", category, 0).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
