package env

import (
	"slices"
	"testing"
)

func CheckResults(t *testing.T, res any, expected any, err error, shouldBeError bool) {
	if shouldBeError && err == nil {
		t.Error("error should be hit")
	}

	if !shouldBeError && err != nil {
		t.Error("no error should be hit")
	}

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func CheckSliceResults[T comparable](t *testing.T, res []T, expected []T, err error, shouldBeError bool) {
	if shouldBeError && err == nil {
		t.Error("error should be hit")
	}

	if !shouldBeError && err != nil {
		t.Error("no error should be hit")
	}

	if !slices.Equal(res, expected) {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}
