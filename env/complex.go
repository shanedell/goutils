package env

import (
	"os"
	"strconv"
)

// GetComplex64 looks up the environment variable and returns an complex64 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex64, the default value and error is returned..
// If the variable is set, and can be parsed to an complex64, the complex64 value is returned.
func GetComplex64(varName string, defaultValue complex64) (complex64, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	complex64Value, err := strconv.ParseComplex(value, 64)
	if err != nil {
		return defaultValue, err
	}

	return complex64(complex64Value), nil
}

// GetComplex128 looks up the environment variable and returns an complex128 or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but can't be parsed to an complex128, the default value and error is returned.
// If the variable is set, and can be parsed to an complex128, the float64 value is returned.
func GetComplex128(varName string, defaultValue complex128) (complex128, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	complex64Value, err := strconv.ParseComplex(value, 64)
	if err != nil {
		return defaultValue, err
	}

	return complex64Value, nil
}
