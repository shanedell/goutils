package env

import (
	"os"
	"testing"
)

var floatVarName = "TEST_GET_FLOAT"

func float32Test(args *float32TestArgs) {
	os.Setenv(floatVarName, args.Common.VarValue)
	res, err := GetFloat32(floatVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func float64Test(args *float64TestArgs) {
	os.Setenv(floatVarName, args.Common.VarValue)
	res, err := GetFloat64(floatVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func TestGetFloat32DefaultNoError(t *testing.T) {
	args := &float32TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     0.0,
		DefaultValue: 0.0,
	}
	float32Test(args)
}

func TestGetFloat32DefaultError(t *testing.T) {
	args := &float32TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     0.0,
		DefaultValue: 0.0,
	}
	float32Test(args)
}

func TestGetFloat32(t *testing.T) {
	args := &float32TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "100.0",
			ShouldBeError: false,
		},
		Expected:     100.0,
		DefaultValue: 0.0,
	}
	float32Test(args)
}

func TestGetFloat64DefaultNoError(t *testing.T) {
	args := &float64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     0.0,
		DefaultValue: 0.0,
	}
	float64Test(args)
}

func TestGetFloat64DefaultError(t *testing.T) {
	args := &float64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     0.0,
		DefaultValue: 0.0,
	}
	float64Test(args)
}

func TestGetFloat64(t *testing.T) {
	args := &float64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "100.0",
			ShouldBeError: false,
		},
		Expected:     100.0,
		DefaultValue: 0.0,
	}
	float64Test(args)
}
