package env

import (
	"os"
	"strconv"

	"github.com/shanedell/goutils/slice"
)

// GetStringSlice looks up the environment variable and returns an string slice.
// If the looked up variable is not set the default value is returned.
// If the variable is set, its content is returned as an string slice.
func GetStringSlice(varName string, delimiter string, defaultValue []string) []string {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue
	}

	return slice.ConvertStringToSlice(delimiter, value)
}

// GetBoolSlice looks up the environment variable and returns an bool slice or error.
// If the looked up variable is not set the default value is returned.
// If the variable is set, but a value can't be parsed to an bool, the default value and error is returned.
// If the variable is set, and all values can be parsed to an bool, the bool slice is returned.
func GetBoolSlice(varName string, delimiter string, defaultValue []bool) ([]bool, error) {
	value := os.Getenv(varName)
	if len(value) == 0 {
		return defaultValue, nil
	}

	stringSlice := GetStringSlice(varName, delimiter, []string{})
	boolSlice := []bool{}

	for _, item := range stringSlice {
		boolValue, err := strconv.ParseBool(item)
		if err != nil {
			return defaultValue, err
		}
		boolSlice = append(boolSlice, boolValue)
	}

	return boolSlice, nil
}
