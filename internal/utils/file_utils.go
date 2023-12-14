package utils

import (
	"os"
)

// GetEnv returns the value of the specified environment variable, or the fallback if not present
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
