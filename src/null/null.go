package NullComplex

import (
	Complex "github.com/go-composites/complex/src"
	MethodNotImplementedError "github.com/go-composites/error/src/method_not_implemented"
	Result "github.com/go-composites/result/src"
)

/*
NullComplex is the Null-Object variant of Complex.

It satisfies Complex.Interface so callers never have to test for a bare nil:
its value is zero, its arithmetic yields a Result carrying a
"method not implemented" Error, and IsNull() returns true.
*/
type Interface interface {
	Complex.Interface
}

type data struct{}

/*
New returns a NullComplex.
*/
func New() Interface {
	return &data{}
}

func (d data) Real() float64 {
	return 0
}

func (d data) Imaginary() float64 {
	return 0
}

func (d data) ToGoString() string {
	return ``
}

func (d data) IsNull() bool {
	return true
}

func notImplemented(methodName string) Result.Interface {
	return Result.New(
		Result.WithError(
			MethodNotImplementedError.New(methodName),
		),
	)
}

func (d data) Add(Complex.Interface) Result.Interface {
	return notImplemented(`Add`)
}

func (d data) Sub(Complex.Interface) Result.Interface {
	return notImplemented(`Sub`)
}

func (d data) Mul(Complex.Interface) Result.Interface {
	return notImplemented(`Mul`)
}

func (d data) Div(Complex.Interface) Result.Interface {
	return notImplemented(`Div`)
}

func (d data) Abs() float64 {
	return 0
}

func (d data) Conjugate() Complex.Interface {
	return New()
}

func (d data) Equal(Complex.Interface) bool {
	return false
}
