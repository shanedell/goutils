package env

import (
	"log"
	"os"
	"strconv"
)

// helper function to parse bool
func parseBool(varName string, defaultValue bool) (bool, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue, err
	}

	return boolValue, nil
}

// GetBool looks up the environment variable and returns an bool or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an bool, the default value and error is returned.
// If the variable is set, and can be parsed to an bool, the bool value is returned.
func GetBool(varName string, defaultValue bool) (bool, error) {
	return parseBool(varName, defaultValue)
}

// GetBoolNoError looks up the environment variable and returns an bool.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an bool, the default value.
// If the variable is set, and can be parsed to an bool, the bool value is returned.
func GetBoolNoError(varName string, defaultValue bool) bool {
	value, err := parseBool(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to bool: %s", err)
	}
	return value
}
