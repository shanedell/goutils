package env

import (
	"log"
	"os"
	"strconv"
)

// helper function for parsing int
func parseInt(varName string, defaultValue int) (int, error) {
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

// GetInt looks up the environment variable and returns an int and error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an int, the default value and error is returned.
// If the variable is set, and can be parsed to an int, the int value is returned.
func GetInt(varName string, defaultValue int) (int, error) {
	return parseInt(varName, defaultValue)
}

// GetIntNoError looks up the environment variable and returns an int.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an int, the default value is returned.
// If the variable is set, and can be parsed to an int, the int value is returned.
func GetIntNoError(varName string, defaultValue int) int {
	value, err := parseInt(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to int: %s", err)
	}
	return value
}
