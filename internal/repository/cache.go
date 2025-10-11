package repository

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		rdb: rdb,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, dur time.Duration) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "cache.Set").Logger()
	rawB, err := json.Marshal(value)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}

	err = c.rdb.Set(ctx, key, rawB, dur).Err()
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	return nil

}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "cache.Get").Logger()
	result, err := c.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		logger.Error().Err(err).Send()
		return "", err
	}
	return result, nil
}
