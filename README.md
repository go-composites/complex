<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/complex" width="720"></p>

# complex

[![ci](https://github.com/go-composites/complex/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/complex/actions/workflows/ci.yml)

A complex-number composite for **Composition-Oriented Programming**. A `Complex`
models Ruby's `Complex` (`a + bi`) over a Go `complex128` value and exposes its
arithmetic as **fallible operations that return a `Result`** — so failures (the
canonical example being a division by `0+0i`) are *values*, never panics and
never `nil`.

```golang
quotient := numerator.Div(denominator)
if quotient.HasError() {
    fmt.Println(quotient.Error().Message()) // "division by zero"
} else {
    fmt.Println(quotient.Payload().(Complex.Interface).ToGoString())
}
```

`Complex` follows the org's Null-Object / never-nil invariant (enforced by the
`nonnil` CI analyzer): the `NullComplex` variant in `src/null` satisfies the same
`Interface` and reports `IsNull() == true`.

> The package is named `Complex` (capitalised) because the lowercase `complex`
> is a Go builtin function and cannot be used as a package name.

## Install

```bash
export GOPRIVATE=github.com/go-composites GOPROXY=direct GOSUMDB=off
go get github.com/go-composites/complex@main
```

## Usage

> [!NOTE] main.go

```golang
package main

import (
    "fmt"

    Complex "github.com/go-composites/complex/src"
)

func main() {
    a := Complex.New(1, 2) // 1+2i
    b := Complex.New(3, 4) // 3+4i

    // Arithmetic returns a Result.
    product := a.Mul(b)
    fmt.Println(product.Payload().(Complex.Interface).ToGoString()) // -5+10i

    // Division by 0+0i is a value, not a NaN/Inf or a panic.
    div := a.Div(Complex.New(0, 0))
    fmt.Println("has error:", div.HasError())  // true
    fmt.Println(div.Error().Message())         // division by zero

    fmt.Println(b.Abs())                        // 5
    fmt.Println(b.Conjugate().ToGoString())     // 3-4i
}
```

```bash
$ task build
```

## API

Constructors

- `New(real, imag float64) Interface` — build `a + bi` from its parts.
- `FromReal(r float64) Interface` — a Complex with a zero imaginary part.
- `null.New() Interface` — the `NullComplex` Null-Object (`IsNull() == true`).

Accessors

- `Real() float64`, `Imaginary() float64`.
- `ToGoString() string` — `"a+bi"` / `"a-bi"`.
- `IsNull() bool`.

Arithmetic (each returns `Result.Interface`)

- `Add(other) Result` / `Sub(other) Result` / `Mul(other) Result`.
- `Div(other) Result` — a `Result` carrying `Error.New("division by zero")`
  when `other` is `0+0i`.

Magnitude, conjugate and equality

- `Abs() float64` — the magnitude (modulus) via `math/cmplx`.
- `Conjugate() Interface` — `a - bi`.
- `Equal(other) bool`.

## License

BSD-3-Clause — see [LICENSE](./LICENSE).
