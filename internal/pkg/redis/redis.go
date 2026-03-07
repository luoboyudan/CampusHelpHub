package redis

import (
	"campushelphub/internal/config"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	RedisClient *redis.Client
}

func NewRedisService(config *config.Config) *RedisService {
	return &RedisService{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:         config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port),
			Password:     config.Redis.Password,
			DB:           config.Redis.DB,
			PoolSize:     config.Redis.PoolSize,
			MinIdleConns: config.Redis.MinIdleConns,
			MaxRetries:   config.Redis.MaxRetries,
			DialTimeout:  time.Duration(config.Redis.DialTimeout) * time.Second,
			ReadTimeout:  time.Duration(config.Redis.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(config.Redis.WriteTimeout) * time.Second,
		}),
	}
}

func (r *RedisService) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return r.RedisClient.Set(ctx, key, value, expire).Err()
}

func (r *RedisService) Get(ctx context.Context, key string) (interface{}, error) {
	return r.RedisClient.Get(ctx, key).Result()
}
