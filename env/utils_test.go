package env

import "testing"

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
