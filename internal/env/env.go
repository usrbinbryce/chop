package env

import (
	"errors"
	"os"
)

var (
	ErrEnvVarNotFound = errors.New("environment variable not found")
)

// GetString looks up an specified environment variable and returns
// it if it exists, or the provided fallback if it doesn't.
//
// For retrieving a required environment variable, use GetRequiredString.
func GetString(key, fallback string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}

	return val
}

// GetRequiredString is the same as GetString, except
// GetRequiredString returns an error if the specified
// environment variable does not exist.
func GetRequiredString(key string) (string, error) {
	val, exist := os.LookupEnv(key)
	if !exist {
		return "", ErrEnvVarNotFound
	}

	return val, nil
}
