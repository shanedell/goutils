package env

import "os"

// GetString looks up the environment variable and returns an string.
// If the looked up variable is not set the default value is returned.
// If the variable is set, its content is returned as an string.
func GetString(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
