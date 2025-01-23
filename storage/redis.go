package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
	"fmt"
)

var rdb *redis.Client

func init() {
	// Conecta a Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Dirección de Redis
		Password: "",               // Sin contraseña
		DB:       0,                // Base de datos por defecto
	})

	// Verifica la conexión
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a Redis")
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