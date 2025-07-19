package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client
var redisAddr string
var redisPassword string
var redisDB int

type RedisClient struct {
	client *redis.Client
}

func InitRedis(addr string, password string, db int) error {
	redisAddr = addr
	redisPassword = password
	redisDB = db

	// Don't use sync.Once with prefork mode - each process needs its own connection
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("could not connect to Redis: %v", err)
	}

	// fmt.Print("Redis connected successfully\n")
	return nil
}

func (r *RedisClient) Set(key string, value interface{}, expirationSeconds int) error {
	expiration := time.Duration(expirationSeconds) * time.Second
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not set value in Redis: %v", err)
	}
	return nil
}

func Get(key string) (string, error) {
	if redisClient == nil {
		if redisAddr != "" {
			err := InitRedis(redisAddr, redisPassword, redisDB)
			if err != nil {
				return "", fmt.Errorf("redis client not initialized and re-initialization failed: %v", err)
			}
		} else {
			return "", fmt.Errorf("redis client not initialized. Call InitRedis() first")
		}
	}

	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("could not get value from Redis: %v", err)
	}
	return val, nil
}

func NewRedisClient(addr string, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("could not connect to Redis: %v", err))
	}

	return &RedisClient{client: rdb}
}

func Set(key string, value interface{}, expirationSeconds int) error {
	if redisClient == nil {
		if redisAddr != "" {
			err := InitRedis(redisAddr, redisPassword, redisDB)
			if err != nil {
				return fmt.Errorf("redis client not initialized and re-initialization failed: %v", err)
			}
		} else {
			return fmt.Errorf("redis client not initialized. Call InitRedis() first")
		}
	}

	expiration := time.Duration(expirationSeconds) * time.Second
	err := redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not set value in Redis: %v", err)
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", fmt.Errorf("could not get value from Redis: %v", err)
	}
	return val, nil
}
