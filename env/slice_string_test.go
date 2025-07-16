package env

import (
	"os"
	"testing"
)

var stringSliceVarName = "TEST_GET_STRING_SLICE"

func stringSliceTest(args *sliceTestArgs[string]) {
	os.Setenv(stringSliceVarName, args.Common.VarValue)
	res := GetStringSlice(stringSliceVarName, ",", args.DefaultValue)
	CheckSliceResults(args.Common.T, res, args.Expected, nil, args.Common.ShouldBeError)
}

func TestGetStringSliceDefault(t *testing.T) {
	args := &sliceTestArgs[string]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     []string{"default"},
		DefaultValue: []string{"default"},
	}
	stringSliceTest(args)
}

func TestGetStringSliceNonDefault(t *testing.T) {
	args := &sliceTestArgs[string]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "setvalue",
			ShouldBeError: false,
		},
		Expected:     []string{"setvalue"},
		DefaultValue: []string{},
	}
	stringSliceTest(args)
}

func TestGetStringSliceMultipleValues(t *testing.T) {
	args := &sliceTestArgs[string]{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "value1,value2",
			ShouldBeError: false,
		},
		Expected:     []string{"value1", "value2"},
		DefaultValue: []string{},
	}
	stringSliceTest(args)
}
