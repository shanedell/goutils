package env

import (
	"os"
	"testing"
)

var boolSliceVarName = "TEST_GET_BOOL_SLICE"

func boolSliceTest(args *sliceTestArgs[bool]) {
	os.Setenv(boolSliceVarName, args.Common.VarValue)
	res, err := GetBoolSlice(boolSliceVarName, ",", args.DefaultValue)
	CheckSliceResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func TestGetBoolSliceDefaultFalseNoError(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     []bool{false},
		DefaultValue: []bool{false},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceDefaultFalseError(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     []bool{false},
		DefaultValue: []bool{false},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceDefaultTrueNoError(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     []bool{true},
		DefaultValue: []bool{true},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceDefaultTrueError(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     []bool{true},
		DefaultValue: []bool{true},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceTrue(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "true",
			ShouldBeError: false,
		},
		Expected:     []bool{true},
		DefaultValue: []bool{false},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceFalse(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "false",
			ShouldBeError: false,
		},
		Expected:     []bool{false},
		DefaultValue: []bool{true},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceTrueFalse(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "true,false",
			ShouldBeError: false,
		},
		Expected:     []bool{true, false},
		DefaultValue: []bool{false, true},
	}
	boolSliceTest(args)
}

func TestGetBoolSliceFalseTrue(t *testing.T) {
	args := &sliceTestArgs[bool]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "false,true",
			ShouldBeError: false,
		},
		Expected:     []bool{false, true},
		DefaultValue: []bool{true, false},
	}
	boolSliceTest(args)
}
