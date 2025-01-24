package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"url-shortener/config"
)

var rdb *redis.Client

func InitRedis(cfg *config.Config) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	return rdb.Ping(context.Background()).Err()
}

// CacheURL almacena una URL en Redis
func CacheURL(shortID, longURL string) error {
	err := rdb.Set(context.Background(), shortID, longURL, 24*time.Hour).Err()
	return err
}

// GetCachedURL obtiene una URL desde Redis
func GetCachedURL(shortID string) (string, error) {
	longURL, err := rdb.Get(context.Background(), shortID).Result()
	if err == redis.Nil {
		return "", nil // No se encontró en Redis
	}
	return longURL, err
}