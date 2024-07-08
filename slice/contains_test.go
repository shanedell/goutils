package slice

import (
	"strconv"
	"testing"
)

// BEGIN: ContainsString tests

func TestContainsStringFalse(t *testing.T) {
	expected := false

	list := []string{}

	res := ContainsString(list, strconv.FormatBool(expected))

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsStringTrue(t *testing.T) {
	expected := true

	list := []string{strconv.FormatBool(expected)}

	res := ContainsString(list, strconv.FormatBool(expected))

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsString tests

// BEGIN: ContainsBool tests

func TestContainsBoolFalse(t *testing.T) {
	expected := false

	list := []bool{}

	res := ContainsBool(list, expected)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsBoolTrue(t *testing.T) {
	expected := true

	list := []bool{expected}

	res := ContainsBool(list, expected)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsBool tests

// BEGIN: ContainsInt tests

func TestContainsIntFalse(t *testing.T) {
	expected := false

	list := []int{}

	res := ContainsInt(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsIntTrue(t *testing.T) {
	expected := true

	list := []int{1}

	res := ContainsInt(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsInt tests

// BEGIN: ContainsInt32 tests

func TestContainsInt32False(t *testing.T) {
	expected := false

	list := []int32{}

	res := ContainsInt32(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsInt32True(t *testing.T) {
	expected := true

	list := []int32{1}

	res := ContainsInt32(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsInt32 tests

// BEGIN: ContainsInt64 tests

func TestContainsInt64False(t *testing.T) {
	expected := false

	list := []int64{}

	res := ContainsInt64(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsInt64True(t *testing.T) {
	expected := true

	list := []int64{1}

	res := ContainsInt64(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsInt64 tests

// BEGIN: ContainsFloat32 tests

func TestContainsFloat32False(t *testing.T) {
	expected := false

	list := []float32{}

	res := ContainsFloat32(list, 1.0)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsFloat32True(t *testing.T) {
	expected := true

	list := []float32{1.0}

	res := ContainsFloat32(list, 1.0)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsFloat32 tests

// BEGIN: ContainsFloat64 tests

func TestContainsFloat64False(t *testing.T) {
	expected := false

	list := []float64{}

	res := ContainsFloat64(list, 1.0)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsFloat64True(t *testing.T) {
	expected := true

	list := []float64{1}

	res := ContainsFloat64(list, 1.0)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsFloat64 tests

// BEGIN: ContainsComplex64 tests

func TestContainsComplex64False(t *testing.T) {
	expected := false

	list := []complex64{}

	res := ContainsComplex64(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsComplex64True(t *testing.T) {
	expected := true

	list := []complex64{1}

	res := ContainsComplex64(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsComplex64 tests

// BEGIN: ContainsComplex128 tests

func TestContainsComplex128False(t *testing.T) {
	expected := false

	list := []complex128{}

	res := ContainsComplex128(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

func TestContainsComplex128True(t *testing.T) {
	expected := true

	list := []complex128{1}

	res := ContainsComplex128(list, 1)

	if res != expected {
		t.Errorf("result was incorrect, got: %v, want: %v", res, expected)
	}
}

// END: ContainsComplex64 tests
