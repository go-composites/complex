package main

import (
	"fmt"

	Complex "github.com/go-composites/complex/src"
	Result "github.com/go-composites/result/src"
)

func report(label string, result Result.Interface) {
	if result.HasError() {
		fmt.Printf("%s -> error: %s\n", label, result.Error().Message())
		return
	}
	fmt.Printf("%s -> %s\n", label, result.Payload().(Complex.Interface).ToGoString())
}

func main() {
	a := Complex.New(1, 2)
	b := Complex.New(3, 4)
	zero := Complex.New(0, 0)

	report("(1+2i) * (3+4i)", a.Mul(b))
	fmt.Printf("Abs(3+4i) = %g\n", b.Abs())

	report("(1+2i) + (3+4i)", a.Add(b))
	report("(3+4i) - (1+2i)", b.Sub(a))

	// Division by 0+0i is a value, not a NaN/Inf or a panic.
	report("(1+2i) / (0+0i)", a.Div(zero))

	fmt.Println("conj(3+4i) =", b.Conjugate().ToGoString())
}
