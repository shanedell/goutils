package env // nolint:dupl

import (
	"log"
	"os"
	"strconv"
)

// helper function for parsing float32
func parseFloat32(varName string, defaultValue float32) (float32, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	float64Value, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue, err
	}

	return float32(float64Value), nil
}

// helper function for parsing float64
func parseFloat64(varName string, defaultValue float64) (float64, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	float64Value, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue, err
	}

	return float64Value, nil
}

// GetFloat32 looks up the environment variable and returns an float32 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an float32, the default value and error is returned.
// If the variable is set, and can be parsed to an float32, the float32 value is returned.
func GetFloat32(varName string, defaultValue float32) (float32, error) {
	return parseFloat32(varName, defaultValue)
}

// GetFloat32NoError looks up the environment variable and returns an float32.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an float32, the default value is returned.
// If the variable is set, and can be parsed to an float32, the float32 value is returned.
func GetFloat32NoError(varName string, defaultValue float32) float32 {
	value, err := parseFloat32(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to float32: %s", err)
	}
	return value
}

// GetFloat64 looks up the environment variable and returns an float64 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an float64 the default value and error is returned.
// If the variable is set, and can be parsed to an float64, the float64 value is returned.
func GetFloat64(varName string, defaultValue float64) (float64, error) {
	return parseFloat64(varName, defaultValue)
}

// GetFloat64NoError looks up the environment variable and returns an float64.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an float64 the default value is returned.
// If the variable is set, and can be parsed to an float64, the float64 value is returned.
func GetFloat64NoError(varName string, defaultValue float64) float64 {
	value, err := parseFloat64(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to float64: %s", err)
	}
	return value
}
