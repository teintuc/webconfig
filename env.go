package main

import (
	"os"
)

/* Get the value from the environment or get the default values */
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
