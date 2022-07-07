package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// New returns error if fails to read .env/.env.local file
func New() error {
	paths := []string{}

	if _, err := os.Stat(".env"); err == nil {
		paths = append(paths, ".env")
	}

	if _, err := os.Stat(".env.local"); err == nil {
		paths = append(paths, ".env")
	}

	if err := godotenv.Load(paths...); err != nil {
		return err
	}

	return nil
}

// Get returns the environment variable as a string, or an empty string when undefined.
func Get(key string) string {
	return os.Getenv(key)
}

// GetString returns the environment variable as a string, or the default value when undefined.
func GetString(key, defVal string) string {
	val := Get(key)
	if val == "" {
		return defVal
	}
	return val
}

// GetBool returns the environment variable as a bool, or the default value when undefined or unparsable.
func GetBool(key string, defVal bool) bool {
	val, err := strconv.ParseBool(Get(key))
	if err != nil {
		return defVal
	}
	return val
}

// GetInt returns the environment variable as a int, or the default value when undefined.
func GetInt(key string, defVal int) (int, error) {
	raw := os.Getenv(key)
	if raw == "" {
		return defVal, nil
	}
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, err
	}
	return val, nil
}
