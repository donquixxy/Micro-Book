package utils

import "github.com/google/uuid"

func RandomIDGenerator() string {
	return uuid.NewString()
}
