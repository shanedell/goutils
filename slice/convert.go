package slice

import (
	"slices"
	"strings"
)

// ConvertStringToSlice splits a string by the delimiter and returns the list.
func ConvertStringToSlice(delimiter string, data string) []string {
	return slices.DeleteFunc(
		strings.Split(data, delimiter),
		func(e string) bool {
			return e == "" || e == " "
		},
	)
}
