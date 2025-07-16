package slice

import "slices"

// ContainsString checks if the passed string value is inside of the list of strings.
func ContainsString(list []string, searchValue string) bool {
	return slices.Contains(list, searchValue)
}

// ContainsBool checks if the passed boolean value is inside of the list of booleans.
func ContainsBool(list []bool, searchValue bool) bool {
	return slices.Contains(list, searchValue)
}

// ContainsInt checks if the passed int value is inside of the list of ints.
func ContainsInt(list []int, searchValue int) bool {
	return slices.Contains(list, searchValue)
}

// ContainsInt32 checks if the passed int32 value is inside of the list of int32s.
func ContainsInt32(list []int32, searchValue int32) bool {
	return slices.Contains(list, searchValue)
}

// ContainsInt64 checks if the passed int64 value is inside of the list of int64s.
func ContainsInt64(list []int64, searchValue int64) bool {
	return slices.Contains(list, searchValue)
}

// ContainsFloat32 checks if the passed float32 value is inside of the list of float32s.
func ContainsFloat32(list []float32, searchValue float32) bool {
	return slices.Contains(list, searchValue)
}

// ContainsFloat64 checks if the passed float64 value is inside of the list of float64s.
func ContainsFloat64(list []float64, searchValue float64) bool {
	return slices.Contains(list, searchValue)
}

// ContainsComplex64 checks if the passed complex64 value is inside of the list of complex64s.
func ContainsComplex64(list []complex64, searchValue complex64) bool {
	return slices.Contains(list, searchValue)
}

// ContainsComplex128 checks if the passed complex128 value is inside of the list of complex128s.
func ContainsComplex128(list []complex128, searchValue complex128) bool {
	return slices.Contains(list, searchValue)
}
