package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SongRedis struct {
	db *RedisDB
}

func NewSongRedis(redisDB *RedisDB) *SongRedis {
	return &SongRedis{
		db: redisDB,
	}
}

func (r *SongRedis) CacheUserSong(userID, songID uint, data string) error {
	cacheKey := fmt.Sprintf("song:%d:%d", userID, songID)

	err := r.db.redisClient.Set(context.Background(), cacheKey, data, time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *SongRedis) GetUserCachedSongByID(userID, songID uint) (string, error) {
	cacheKey := fmt.Sprintf("song:%d:%d", userID, songID)
	data, err := r.db.redisClient.Get(context.Background(), cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrCacheNotFound
		}
		return "", err
	}
	return data, nil
}

func (r *SongRedis) DeleteUserCachedSong(userID, songID uint) error {
	cacheKey := fmt.Sprintf("song:%d:%d", userID, songID)

	err := r.db.redisClient.Del(context.Background(), cacheKey).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	return nil
}
