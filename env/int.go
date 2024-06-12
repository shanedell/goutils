package env

import (
	"os"
	"strconv"
)

// GetInt looks up the environment variable and returns an int.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an int, the default value and error is returned.
// If the variable is set, and can be parsed to an int, the int value is returned.
func GetInt(varName string, defaultValue int) (int, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue, err
	}

	return intValue, nil
}
