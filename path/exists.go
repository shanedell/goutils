package path

import (
	"errors"
	"os"
)

// Exists check if the path provided exists.
func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
