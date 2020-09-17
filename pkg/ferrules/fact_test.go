package ferrules

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vikyd/zero"
)

var _ = Describe("Fact", func() {
	Context("Create Facts", func() {
		set := FactSet()
		in := set.(*factSet)
		Expect(len(in.facts)).To(Equal(0))
	})

	Context("Facts Operation", func() {
		var set Facts

		BeforeEach(func() {
			set = FactSet()
		})

		AfterEach(func() {
			set = nil
		})

		It("add operation", func() {
			var fact *Fact
			in := set.(*factSet)
			Expect(len(in.facts)).To(Equal(0))

			fact = &Fact{"fact A", "A"}
			set.Add(fact)
			Expect(len(in.facts)).To(Equal(1))

			fact = &Fact{"fact B", "B"}
			set.Add(fact)
			Expect(len(in.facts)).To(Equal(2))

			fact = &Fact{"fact A", "AA"}
			set.Add(fact)
			Expect(len(in.facts)).To(Equal(2))
		})

		It("put operation", func() {
			in := set.(*factSet)
			Expect(len(in.facts)).To(Equal(0))

			set.Put("fact A", "A")
			Expect(len(in.facts)).To(Equal(1))

			set.Put("fact B", "B")
			Expect(len(in.facts)).To(Equal(2))

			set.Put("fact A", "AA")
			Expect(len(in.facts)).To(Equal(2))
		})

		It("remove operation", func() {
			in := set.(*factSet)
			Expect(len(in.facts)).To(Equal(0))

			set.Put("fact A", "A")
			Expect(len(in.facts)).To(Equal(1))

			set.Put("fact B", "B")
			Expect(len(in.facts)).To(Equal(2))

			set.Remove("fact A")
			Expect(len(in.facts)).To(Equal(1))

			set.Remove("fact unknown")
			Expect(len(in.facts)).To(Equal(1))
		})

		It("get operation", func() {
			var fact *Fact
			var err error
			in := set.(*factSet)
			Expect(len(in.facts)).To(Equal(0))

			set.Put("fact A", "A")
			Expect(len(in.facts)).To(Equal(1))

			set.Put("fact B", "B")
			Expect(len(in.facts)).To(Equal(2))

			fact, err = set.Get("fact A")
			Expect(string(fact.Name)).To(Equal("fact A"))
			Expect(fact.Value.(string)).To(Equal("A"))
			Expect(err).ShouldNot(HaveOccurred())

			set.Put("fact A", "AA")
			fact, err = set.Get("fact A")
			Expect(string(fact.Name)).To(Equal("fact A"))
			Expect(fact.Value.(string)).To(Equal("AA"))
			Expect(err).ShouldNot(HaveOccurred())

			fact = &Fact{"fact A", "A"}
			set.Add(fact)
			fact, err = set.Get("fact A")
			Expect(string(fact.Name)).To(Equal("fact A"))
			Expect(fact.Value.(string)).To(Equal("A"))
			Expect(err).ShouldNot(HaveOccurred())

			set.Remove("fact A")
			fact, err = set.Get("fact A")
			Expect(zero.IsZeroVal(fact)).Should(BeTrue())
			Expect(err).Should(HaveOccurred())
		})
	})
})
