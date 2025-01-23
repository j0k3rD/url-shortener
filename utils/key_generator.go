package utils

import (
	"math/rand"
	"time"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength   = 7
)

// GenerateRandomKey genera una clave aleatoria de 7 caracteres alfanuméricos
func GenerateRandomKey() string {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, keyLength)
	for i := range key {
		key[i] = base62Chars[rand.Intn(len(base62Chars))]
	}
	return string(key)
}