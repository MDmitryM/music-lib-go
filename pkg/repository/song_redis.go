package repository

import (
	"context"
	"fmt"
	"time"
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
