package ferrules

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rule Engine", func() {
	Context("Test interal fire", func() {
		rules := RuleSet(
			RuleBuilder().
				Given(func(facts Facts) bool {
					return true
				}).
				Will(func(facts Facts) {
					fact, _ := facts.Get("000")
					fact.Value = true
				}).
				Priority(1).
				WithName("Rule 000").
				WithDescription("Always say hello world").
				Build(),

			RuleBuilder().
				Given(func(facts Facts) bool {
					return false
				}).
				Will(func(facts Facts) {
					fact, _ := facts.Get("001")
					fact.Value = true
				}).
				Priority(1).
				WithName("Rule 001").
				WithDescription("Always say hello world").
				Build(),

			RuleBuilder().
				Given(func(facts Facts) bool {
					return true
				}).
				Will(func(facts Facts) {
					fact, _ := facts.Get("002")
					fact.Value = true
				}).
				Priority(1).
				WithName("Rule 002").
				WithDescription("Always say hello world").
				Build(),
		)

		It("fire", func() {
			engine := &DefaultRuleEngine{}

			f0 := Fact{"000", false}
			f1 := Fact{"001", false}
			f2 := Fact{"002", false}
			facts := FactSet()
			facts.Add(&f0)
			facts.Add(&f1)
			facts.Add(&f2)

			fire(engine, rules.values(), facts)
			Expect(f0.Value).Should(BeTrue())
			Expect(f1.Value).Should(BeFalse())
			Expect(f2.Value).Should(BeTrue())
		})
	})
})
