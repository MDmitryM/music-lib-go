package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RedisDB struct {
	redisClient *redis.Client
}

func NewRedisDB(cfg RedisConfig) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed, %s", err.Error())
	}
	logrus.Println("redis connected successfully")
	return &RedisDB{
		redisClient: client,
	}, nil
}
