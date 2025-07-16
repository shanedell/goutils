package env

import "testing"

type commonTestArgs struct {
	T             *testing.T
	VarValue      string
	ShouldBeError bool
}

type boolTestArgs struct {
	Common       *commonTestArgs
	Expected     bool
	DefaultValue bool
}

type complex64TestArgs struct {
	Common       *commonTestArgs
	Expected     complex64
	DefaultValue complex64
}

type complex128TestArgs struct {
	Common       *commonTestArgs
	Expected     complex128
	DefaultValue complex128
}

type float32TestArgs struct {
	Common       *commonTestArgs
	Expected     float32
	DefaultValue float32
}

type float64TestArgs struct {
	Common       *commonTestArgs
	Expected     float64
	DefaultValue float64
}

type intTestArgs struct {
	Common       *commonTestArgs
	Expected     int
	DefaultValue int
}

type stringTestArgs struct {
	Common       *commonTestArgs
	Expected     string
	DefaultValue string
}

type sliceTestArgs[T comparable] struct {
	Common       *commonTestArgs
	Expected     []T
	DefaultValue []T
}
