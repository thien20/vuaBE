package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCacheFromClient(redisClient *redis.Client) *Cache {
	return &Cache{client: redisClient}
}

// Set stores a key-value pair in Redis with an optional expiration time
func (c *Cache) SetCache(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}
	return nil
}

// Get retrieves the value associated with the given key
func (c *Cache) GetCache(key string) (string, error) {
	ctx := context.Background()
	value, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", fmt.Errorf("failed to get key %s: %w", key, err)
	}
	return value, nil
}

func (c *Cache) DeleteCache(key string) error {
	ctx := context.Background()
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %w", key, err)
	}
	return nil
}
