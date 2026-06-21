package NullComplex_test

import (
	Complex "github.com/go-composites/complex/src"
	NullComplex "github.com/go-composites/complex/src/null"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("NullComplex", func() {
	var n NullComplex.Interface
	ginkgo.BeforeEach(func() {
		n = NullComplex.New()
	})

	ginkgo.It("satisfies the Complex interface", func() {
		var _ Complex.Interface = n
	})
	ginkgo.It("reports IsNull() true", func() {
		gomega.Expect(n.IsNull()).To(gomega.BeTrue())
	})
	ginkgo.It("converts to zero values", func() {
		gomega.Expect(n.Real()).To(gomega.BeEquivalentTo(0.0))
		gomega.Expect(n.Imaginary()).To(gomega.BeEquivalentTo(0.0))
		gomega.Expect(n.ToGoString()).To(gomega.Equal(``))
	})

	ginkgo.It("Add returns an error result", func() {
		r := n.Add(Complex.New(1, 1))
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Add"))
	})
	ginkgo.It("Sub returns an error result", func() {
		r := n.Sub(Complex.New(1, 1))
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Sub"))
	})
	ginkgo.It("Mul returns an error result", func() {
		r := n.Mul(Complex.New(1, 1))
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Mul"))
	})
	ginkgo.It("Div returns an error result", func() {
		r := n.Div(Complex.New(1, 1))
		gomega.Expect(r.HasError()).To(gomega.BeTrue())
		gomega.Expect(r.Error().Message()).To(gomega.ContainSubstring("Div"))
	})
	ginkgo.It("Abs returns zero", func() {
		gomega.Expect(n.Abs()).To(gomega.BeEquivalentTo(0.0))
	})
	ginkgo.It("Conjugate returns the null complex", func() {
		gomega.Expect(n.Conjugate().IsNull()).To(gomega.BeTrue())
	})
	ginkgo.It("Equal is always false", func() {
		gomega.Expect(n.Equal(NullComplex.New())).To(gomega.BeFalse())
		gomega.Expect(n.Equal(Complex.New(0, 0))).To(gomega.BeFalse())
	})
})
