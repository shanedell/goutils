package path

import (
	"testing"
)

func TestExistsFalse(t *testing.T) {
	expected := false

	res := Exists("BAD_PATH")

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestExistsTrue(t *testing.T) {
	expected := false

	res := Exists("path/exists.go")

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}
