package env

import (
	"os"
	"testing"
)

var intVarName = "TEST_GET_INT"

func intTest(args *intTestArgs) {
	os.Setenv(intVarName, args.Common.VarValue)
	res, err := GetInt(intVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func TestGetIntDefaultNoError(t *testing.T) {
	args := &intTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     0,
		DefaultValue: 0,
	}
	intTest(args)
}

func TestGetIntDefaultError(t *testing.T) {
	args := &intTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     0,
		DefaultValue: 0,
	}
	intTest(args)
}

func TestGetInt(t *testing.T) {
	args := &intTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "100",
			ShouldBeError: false,
		},
		Expected:     100,
		DefaultValue: 0,
	}
	intTest(args)
}
