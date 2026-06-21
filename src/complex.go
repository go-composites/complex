package Complex

import (
	"math/cmplx"
	"strconv"

	Error "github.com/go-composites/error/src"
	Result "github.com/go-composites/result/src"
)

/*
Complex is a complex-number composite over a Go complex128 value, modelling
Ruby's Complex (a + bi).

Its fallible operations (notably Div) return a Result.Interface so that
failures — such as a division by zero — are values rather than panics.
*/
type Interface interface {
	Real() float64
	Imaginary() float64
	ToGoString() string
	IsNull() bool
	Add(Interface) Result.Interface
	Sub(Interface) Result.Interface
	Mul(Interface) Result.Interface
	Div(Interface) Result.Interface
	Abs() float64
	Conjugate() Interface
	Equal(Interface) bool
}

type data struct {
	value complex128
}

/*
New is the Complex constructor: it builds a + bi from its real and imaginary
parts.

	c := Complex.New(3, 4) // 3+4i
*/
func New(real, imag float64) Interface {
	return &data{value: complex(real, imag)}
}

/*
FromReal builds a Complex whose imaginary part is zero.

	c := Complex.FromReal(3) // 3+0i
*/
func FromReal(r float64) Interface {
	return &data{value: complex(r, 0)}
}

/*
Real returns the real part of the Complex.
*/
func (d data) Real() float64 {
	return real(d.value)
}

/*
Imaginary returns the imaginary part of the Complex.
*/
func (d data) Imaginary() float64 {
	return imag(d.value)
}

/*
ToGoString returns the textual representation of the Complex as "a+bi" or
"a-bi", choosing the sign of the imaginary part sensibly.
*/
func (d data) ToGoString() string {
	re := strconv.FormatFloat(real(d.value), 'g', -1, 64)
	im := imag(d.value)
	sign := "+"
	if im < 0 {
		sign = "-"
		im = -im
	}
	return re + sign + strconv.FormatFloat(im, 'g', -1, 64) + "i"
}

/*
IsNull reports whether the Complex is the Null-Object variant.

A concrete Complex is never null.
*/
func (d data) IsNull() bool {
	return false
}

/*
Add returns a Result whose payload is the sum of the receiver and other.
*/
func (d data) Add(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(
			New(d.Real()+other.Real(), d.Imaginary()+other.Imaginary()),
		),
	)
}

/*
Sub returns a Result whose payload is the difference of the receiver and other.
*/
func (d data) Sub(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(
			New(d.Real()-other.Real(), d.Imaginary()-other.Imaginary()),
		),
	)
}

/*
Mul returns a Result whose payload is the product of the receiver and other,
following the rule (a+bi)(c+di) = (ac-bd) + (ad+bc)i.
*/
func (d data) Mul(other Interface) Result.Interface {
	a, b := d.Real(), d.Imaginary()
	c, e := other.Real(), other.Imaginary()
	return Result.New(
		Result.WithPayload(
			New(a*c-b*e, a*e+b*c),
		),
	)
}

/*
Div returns a Result whose payload is the quotient of the receiver and other.

When other is 0+0i the Result carries an Error ("division by zero") instead of a
payload — the division never panics, never yields NaN/Inf, and never returns
nil.
*/
func (d data) Div(other Interface) Result.Interface {
	if other.Real() == 0 && other.Imaginary() == 0 {
		return Result.New(
			Result.WithError(
				Error.New("division by zero"),
			),
		)
	}
	q := d.value / complex(other.Real(), other.Imaginary())
	return Result.New(
		Result.WithPayload(
			New(real(q), imag(q)),
		),
	)
}

/*
Abs returns the magnitude (modulus) of the Complex, sqrt(a*a + b*b).
*/
func (d data) Abs() float64 {
	return cmplx.Abs(d.value)
}

/*
Conjugate returns the complex conjugate a-bi of the receiver.
*/
func (d data) Conjugate() Interface {
	return New(d.Real(), -d.Imaginary())
}

/*
Equal reports whether the receiver and other hold the same complex value.
*/
func (d data) Equal(other Interface) bool {
	return d.Real() == other.Real() && d.Imaginary() == other.Imaginary()
}
