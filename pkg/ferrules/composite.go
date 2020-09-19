package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
	"github.com/vikyd/zero"
)

// AnyOfRules is an OR logical group of rules that
// iterates every rule, applies the first applicable rule
// and ignores the rest of rules.
type AnyOfRules struct {
	name        RuleName
	description string
	priority    RulePriority
	rules       []Rule
	applicable  Rule
}

// Name gives the name of rule
func (g *AnyOfRules) Name() RuleName {
	return g.name
}

// Description gives the description of rule
func (g *AnyOfRules) Description() string {
	return g.description
}

// Priority gives the priority of rule
func (g *AnyOfRules) Priority() RulePriority {
	return g.priority
}

func (g *AnyOfRules) evaluate(facts Facts) bool {
	for _, r := range g.rules {
		if r.evaluate(facts) {
			g.applicable = r
			return true
		}
	}
	return false
}

func (g *AnyOfRules) execute(facts Facts) {
	if !zero.IsZeroVal(g.applicable) {
		g.applicable.execute(facts)
	}
}

type compositeRuleBuilder struct {
	name        RuleName
	description string
	priority    RulePriority
	rules       []Rule
}

func (b *compositeRuleBuilder) AnyOf(rules ...Rule) PriorityOngoing {
	b.rules = rules
	return b
}

func (b *compositeRuleBuilder) Priority(n int) Outgoing {
	b.priority = RulePriority(n)
	return b
}

func (b *compositeRuleBuilder) WithName(name string) Outgoing {
	b.name = RuleName(name)
	return b
}

func (b *compositeRuleBuilder) WithDescription(desc string) Outgoing {
	b.description = desc
	return b
}

func (b *compositeRuleBuilder) Build() Rule {
	defer func() {
		b = nil
	}()

	r := AnyOfRules{}

	if zero.IsZeroVal(b.name) {
		panic("rule name is missing")
	} else {
		r.name = b.name
	}

	if zero.IsZeroVal(b.description) {
		panic("rule description is missing")
	} else {
		r.description = b.description
	}

	r.priority = b.priority

	if zero.IsZeroVal(b.rules) {
		panic("rule condition is missing")
	} else {
		r.rules = b.rules
	}
	return &r
}
