package ferrules_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/u2386/ferrules/pkg/ferrules"
)

var _ = Describe("Listener", func() {
	Context("Rule Engine Listeners", func() {
		called := false

		rules := RuleSet(
			RuleBuilder().
				Given(func(facts Facts) bool {
					return true
				}).
				Will(func(facts Facts) {}).
				WithName("Foo").
				WithDescription("Foo").
				Build(),
		)

		BeforeEach(func() {
			called = false
		})

		It("trigger before evaluate", func() {
			Expect(called).Should(BeFalse())

			engine := &DefaultRuleEngine{
				Listener: []RuleEngineListener{
					{
						Before: func(rules Rules, facts Facts) {
							called = true
						},
					},
				},
			}
			facts := FactSet()
			engine.Fire(rules, facts)

			Expect(called).Should(BeTrue())
		})

		It("trigger after evaluate", func() {
			Expect(called).Should(BeFalse())

			engine := &DefaultRuleEngine{
				Listener: []RuleEngineListener{
					{
						After: func(rules Rules, facts Facts) {
							called = true
						},
					},
				},
			}
			facts := FactSet()
			engine.Fire(rules, facts)

			Expect(called).Should(BeTrue())
		})
	})

})
