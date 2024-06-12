package env

import (
	"os"
	"testing"
)

var boolVarName = "TEST_GET_BOOL"

func boolTest(args *boolTestArgs) {
	os.Setenv(boolVarName, args.Common.VarValue)
	res, err := GetBool(boolVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func TestGetBoolDefaultFalseNoError(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     false,
		DefaultValue: false,
	}
	boolTest(args)
}

func TestGetBoolDefaultFalseError(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     false,
		DefaultValue: false,
	}
	boolTest(args)
}

func TestGetBoolDefaultTrueNoError(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     true,
		DefaultValue: true,
	}
	boolTest(args)
}

func TestGetBoolDefaultTrueError(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     true,
		DefaultValue: true,
	}
	boolTest(args)
}

func TestGetBoolTrue(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "true",
			ShouldBeError: false,
		},
		Expected:     true,
		DefaultValue: false,
	}
	boolTest(args)
}

func TestGetBoolFalse(t *testing.T) {
	args := &boolTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "false",
			ShouldBeError: false,
		},
		Expected:     false,
		DefaultValue: true,
	}
	boolTest(args)
}
