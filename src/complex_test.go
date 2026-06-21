package Complex_test

import (
	Complex "github.com/go-composites/complex/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Complex", func() {

	ginkgo.Describe("constructors", func() {
		ginkgo.It("builds a+bi from its parts", func() {
			c := Complex.New(3, 4)
			gomega.Expect(c.Real()).To(gomega.BeEquivalentTo(3))
			gomega.Expect(c.Imaginary()).To(gomega.BeEquivalentTo(4))
		})
		ginkgo.It("builds a real number with zero imaginary part", func() {
			c := Complex.FromReal(3)
			gomega.Expect(c.Real()).To(gomega.BeEquivalentTo(3))
			gomega.Expect(c.Imaginary()).To(gomega.BeEquivalentTo(0))
		})
		ginkgo.It("is never null", func() {
			gomega.Expect(Complex.New(0, 0).IsNull()).To(gomega.BeFalse())
		})
	})

	ginkgo.Describe("rendering", func() {
		ginkgo.It("renders a positive imaginary part with a plus sign", func() {
			gomega.Expect(Complex.New(3, 4).ToGoString()).To(gomega.Equal("3+4i"))
		})
		ginkgo.It("renders a negative imaginary part with a minus sign", func() {
			gomega.Expect(Complex.New(3, -4).ToGoString()).To(gomega.Equal("3-4i"))
		})
	})

	ginkgo.Describe("arithmetic", func() {
		var a = Complex.New(1, 2)
		var b = Complex.New(3, 4)

		ginkgo.It("adds two complex numbers", func() {
			r := a.Add(b)
			gomega.Expect(r.HasError()).To(gomega.BeFalse())
			gomega.Expect(r.Payload().(Complex.Interface).ToGoString()).To(gomega.Equal("4+6i"))
		})
		ginkgo.It("subtracts two complex numbers", func() {
			r := b.Sub(a)
			gomega.Expect(r.Payload().(Complex.Interface).ToGoString()).To(gomega.Equal("2+2i"))
		})
		ginkgo.It("multiplies two complex numbers", func() {
			r := a.Mul(b)
			gomega.Expect(r.HasError()).To(gomega.BeFalse())
			gomega.Expect(r.Payload().(Complex.Interface).ToGoString()).To(gomega.Equal("-5+10i"))
		})
		ginkgo.It("divides two complex numbers", func() {
			// (1+2i) / (1+0i) = 1+2i
			r := a.Div(Complex.New(1, 0))
			gomega.Expect(r.HasError()).To(gomega.BeFalse())
			gomega.Expect(r.Payload().(Complex.Interface).ToGoString()).To(gomega.Equal("1+2i"))
		})

		ginkgo.Describe("division by zero", func() {
			ginkgo.It("returns a Result carrying an error instead of NaN/Inf", func() {
				r := a.Div(Complex.New(0, 0))
				gomega.Expect(r.HasError()).To(gomega.BeTrue())
				gomega.Expect(r.Error().Message()).To(gomega.Equal("division by zero"))
			})
		})
	})

	ginkgo.Describe("magnitude and conjugate", func() {
		ginkgo.It("computes the absolute value", func() {
			gomega.Expect(Complex.New(3, 4).Abs()).To(gomega.BeEquivalentTo(5.0))
		})
		ginkgo.It("computes the conjugate", func() {
			gomega.Expect(Complex.New(3, 4).Conjugate().ToGoString()).To(gomega.Equal("3-4i"))
		})
	})

	ginkgo.Describe("equality", func() {
		ginkgo.It("reports equality", func() {
			gomega.Expect(Complex.New(3, 4).Equal(Complex.New(3, 4))).To(gomega.BeTrue())
			gomega.Expect(Complex.New(3, 4).Equal(Complex.New(3, 5))).To(gomega.BeFalse())
		})
	})
})
