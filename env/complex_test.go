package env

import (
	"os"
	"testing"
)

var complexVarName = "TEST_GET_COMPLEX"

func complex64Test(args *complex64TestArgs) {
	os.Setenv(complexVarName, args.Common.VarValue)
	res, err := GetComplex64(complexVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func complex128Test(args *complex128TestArgs) {
	os.Setenv(complexVarName, args.Common.VarValue)
	res, err := GetComplex128(complexVarName, args.DefaultValue)
	CheckResults(args.Common.T, res, args.Expected, err, args.Common.ShouldBeError)
}

func TestGetComplex64DefaultNoError(t *testing.T) {
	args := &complex64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     complex64(0.0),
		DefaultValue: complex64(0.0),
	}
	complex64Test(args)
}

func TestGetComplex64DefaultError(t *testing.T) {
	args := &complex64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     complex64(0.0),
		DefaultValue: complex64(0.0),
	}
	complex64Test(args)
}

func TestGetComplex64(t *testing.T) {
	args := &complex64TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "100.0",
			ShouldBeError: false,
		},
		Expected:     complex64(100.0),
		DefaultValue: complex64(0.0),
	}
	complex64Test(args)
}

func TestGetComplex128DefaultNoError(t *testing.T) {
	args := &complex128TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "",
			ShouldBeError: false,
		},
		Expected:     complex128(0.0),
		DefaultValue: complex128(0.0),
	}
	complex128Test(args)
}

func TestGetComplex128DefaultError(t *testing.T) {
	args := &complex128TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "NOT_PARSIBLE",
			ShouldBeError: true,
		},
		Expected:     complex128(0.0),
		DefaultValue: complex128(0.0),
	}
	complex128Test(args)
}

func TestGetComplex128(t *testing.T) {
	args := &complex128TestArgs{
		Common: &commonTestArgs{
			T:             t,
			VarValue:      "100.0",
			ShouldBeError: false,
		},
		Expected:     complex128(100.0),
		DefaultValue: complex128(0.0),
	}
	complex128Test(args)
}
