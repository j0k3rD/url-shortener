package utils

import (
	"github.com/google/uuid"
)

// GenerateUserID genera un UUID único para identificar al usuario
func GenerateUserID() string {
	return uuid.New().String()
}