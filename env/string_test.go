package env

import (
	"os"
	"testing"
)

var stringVarName = "TEST_GET_STRING"

func stringTest(args *stringTestArgs) {
	os.Setenv(stringVarName, args.Common.VarValue)
	res := GetString(stringVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, nil, args.Common.ShouldBeError)
}

func TestGetStringDefault(t *testing.T) {
	args := &stringTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     "default",
		DefaultValue: "default",
	}
	stringTest(args)
}

func TestGetStringNonDefault(t *testing.T) {
	args := &stringTestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "setvalue",
			ShouldBeError: false,
		},
		Expected:     "setvalue",
		DefaultValue: "",
	}
	stringTest(args)
}
