package env

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Load initializes the dotenv files
func Load(fileNames ...string) error {

	return godotenv.Load(fileNames...)
}

// Get returns the value for .env key
func Get(key string, defaultVal string) string {

	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}

	return val
}

// GetInt returns the value for .env key
func GetInt(key string, defaultVal int) int {

	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}

	parsedVal, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("[dotenv] Parse error for key %s, value %s is not a integer\n", key, val)
		return defaultVal
	}

	return parsedVal
}

// GetBool returns the value for .env key
func GetBool(key string, defaultVal bool) bool {

	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}

	val = strings.ToLower(val)
	if val == "true" || val == "yes" || val == "on" {
		return true
	}

	return false
}
