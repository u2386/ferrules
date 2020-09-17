package ferrules

import (
	"fmt"

	. "github.com/u2386/ferrules/internal/types"

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
			Expect(string(rule.Name())).To(Equal("this is a name"))
		})
		It("get description", func() {
			Expect(rule.Description()).To(Equal("this is a description"))
		})
		It("get priority", func() {
			Expect(int(rule.Priority())).To(Equal(1))
		})
	})
})

var _ = Describe("Rule Condition", func() {
	Context("Call rule condition", func() {
		var rule Rule = DefaultRule{
			name:        "this is a name",
			description: "this is a description",
			priority:    1,
			condition:   func(facts Facts) bool { return true },
			actions:     []Action{func(facts Facts) {}},
		}

		It("should be true", func() {
			facts := FactSet()
			Expect(rule.evaluate(facts)).Should(BeTrue())
		})
	})
})

var _ = Describe("Rule Actions", func() {
	Context("Call rule actions", func() {
		var rule Rule = DefaultRule{
			name:        "this is a name",
			description: "this is a description",
			priority:    1,
			condition:   func(facts Facts) bool { return true },
			actions: []Action{func(facts Facts) {
				facts.Put("one", 1)
			}, func(facts Facts) {
				facts.Put("two", 2)
			}},
		}

		It("actions should be called", func() {
			facts := FactSet()
			rule.execute(facts)
			one, _ := facts.Get("one")
			two, _ := facts.Get("two")
			Expect(one.Value.(int)).To(Equal(1))
			Expect(two.Value.(int)).To(Equal(2))
		})
	})
})

var _ = Describe("Rule Build", func() {
	Context("shoud success", func() {
		rule := RuleBuilder().
			Given(func(facts Facts) bool {
				return true
			}).
			Will(func(facts Facts) {
				fmt.Println("Hello")
			}).
			Priority(1).
			WithName("Hello World rule").
			WithDescription("Always say hello world").
			Build()

		Expect(string(rule.Name())).To(Equal("Hello World rule"))
		Expect(rule.Description()).To(Equal("Always say hello world"))
		Expect(int(rule.Priority())).To(Equal(1))
	})
})

var _ = Describe("RuleSet Operations", func() {
	var set Rules
	Context("create a ruleset", func() {
		set = &ruleSet{
			rules: make(map[RuleName]Rule),
		}
		set.Add(DefaultRule{name: "rule 01"})
		set.Add(DefaultRule{name: "rule 02"})
		rules := set.values()
		Expect(len(rules)).To(Equal(2))
	})

	Context("create by RuleSet API", func() {
		set = RuleSet(DefaultRule{name: "rule 01"}, DefaultRule{name: "rule 02"})
		rules := set.values()
		Expect(len(rules)).To(Equal(2))
	})
})
