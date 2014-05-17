package kamino_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/modcloth/kamino"
)

var _ = Describe("NewGenome()", func() {

	var opts map[string]string

	BeforeEach(func() {
		opts = map[string]string{
			"depth":   "50",
			"token":   "abc123",
			"account": "modcloth",
			"repo":    "kamino",
			"cache":   "",
			"ref":     "123",
		}
	})

	It("assigns token from the provided options", func() {
		ret, _ := NewGenome(opts)

		Expect(ret.APIToken).To(Equal("abc123"))
	})

	Context("with a non-integer depth", func() {
		It("returns an error", func() {
			opts["depth"] = "foo"
			ret, err := NewGenome(opts)

			Expect(ret).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with no account specified", func() {
		It("returns an error", func() {
			opts["account"] = ""
			ret, err := NewGenome(opts)

			Expect(ret).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with no repo specified", func() {
		It("returns an error", func() {
			opts["repo"] = ""
			ret, err := NewGenome(opts)

			Expect(ret).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with an empty cache option", func() {
		It("defaults the cache option to `no`", func() {
			ret, _ := NewGenome(opts)

			Expect(ret.UseCache).To(Equal("no"))
		})
	})

	Context("with an invalid cache option", func() {
		It("returns an error", func() {
			opts["cache"] = "foo"
			ret, err := NewGenome(opts)

			Expect(ret).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with no ref specified", func() {
		It("returns an error", func() {
			opts["ref"] = ""
			ret, err := NewGenome(opts)

			Expect(ret).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})
})
