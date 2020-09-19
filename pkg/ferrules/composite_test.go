package ferrules_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/u2386/ferrules/pkg/ferrules"
)

var _ = Describe("Composite", func() {
	var rules Rules
	Context("Create an OR rule group", func() {
		defer GinkgoRecover()

		rules = RuleSet(CompositeRuleBuilder().
			AnyOf(
				RuleBuilder().
					Given(func(facts Facts) bool {
						return false
					}).
					Will(func(facts Facts) {
						fact, _ := facts.Get("Counter")
						v := fact.Value.(int)
						v++
						fact.Value = v
					}).
					WithName("Rule 01").
					WithDescription("Increase 1").
					Build(),

				RuleBuilder().
					Given(func(facts Facts) bool {
						return true
					}).
					Will(func(facts Facts) {
						fact, _ := facts.Get("Counter")
						v := fact.Value.(int)
						v++
						fact.Value = v
					}).
					WithName("Rule 02").
					WithDescription("Increase 1").
					Build(),

				RuleBuilder().
					Given(func(facts Facts) bool {
						return true
					}).
					Will(func(facts Facts) {
						fact, _ := facts.Get("Counter")
						v := fact.Value.(int)
						v++
						fact.Value = v
					}).
					WithName("Rule 03").
					WithDescription("Increase 1").
					Build(),
			).
			WithName("a composite group").
			WithDescription("OR logical").
			Build())

		facts := FactSet()
		facts.Put("Counter", 0)
		engine := DefaultRuleEngine{}
		engine.Fire(rules, facts)

		counter, _ := facts.Get("Counter")
		Expect(counter.Value.(int)).To(Equal(1))
	})
})
