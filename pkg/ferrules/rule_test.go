package ferrules

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rule Creation", func() {
	Context("Create a default rule", func() {
		var rule Rule = DefaultRule{
			name:        "this is a name",
			description: "this is a description",
			priority:    1,
			condition:   func(facts Facts) bool { return true },
			actions:     []Action{func(facts Facts) {}},
		}

		It("get name", func() {
			Expect(rule.Name().String()).To(Equal("this is a name"))
		})
		It("get description", func() {
			Expect(rule.Description()).To(Equal("this is a description"))
		})
		It("get priority", func() {
			Expect(rule.Priority().Int()).To(Equal(1))
		})
	})
})
