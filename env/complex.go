package env // nolint:dupl

import (
	"log"
	"os"
	"strconv"
)

// helper function to parse complex64
func parseComplex64(varName string, defaultValue complex64) (complex64, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	complex128Value, err := strconv.ParseComplex(value, 64)
	if err != nil {
		return defaultValue, err
	}

	return complex64(complex128Value), nil
}

// helper function to parse complex128
func parseComplex128(varName string, defaultValue complex128) (complex128, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	complex128Value, err := strconv.ParseComplex(value, 128)
	if err != nil {
		return defaultValue, err
	}

	return complex128Value, nil
}

// GetComplex64 looks up the environment variable and returns an complex64 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex64, the default value and error is returned..
// If the variable is set, and can be parsed to an complex64, the complex64 value is returned.
func GetComplex64(varName string, defaultValue complex64) (complex64, error) {
	return parseComplex64(varName, defaultValue)
}

// GetComplex64NoError looks up the environment variable and returns an complex64.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex64, the default value is returned..
// If the variable is set, and can be parsed to an complex64, the complex64 value is returned.
func GetComplex64NoError(varName string, defaultValue complex64) complex64 {
	value, err := parseComplex64(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to complex64: %s", err)
	}
	return value
}

// GetComplex128 looks up the environment variable and returns an complex128 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex128, the default value and error is returned.
// If the variable is set, and can be parsed to an complex128, the float64 value is returned.
func GetComplex128(varName string, defaultValue complex128) (complex128, error) {
	return parseComplex128(varName, defaultValue)
}

// GetComplex128NoError looks up the environment variable and returns an complex128.
// Errors are ignored but logged.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex128, the default value is returned.
// If the variable is set, and can be parsed to an complex128, the float64 value is returned.
func GetComplex128NoError(varName string, defaultValue complex128) complex128 {
	value, err := parseComplex128(varName, defaultValue)
	if err != nil {
		log.Printf("error converting value to complex128: %s", err)
	}
	return value
}
