package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
	"github.com/vikyd/zero"
)

type ruleSet struct {
	rules map[RuleName]Rule
}

func (rs *ruleSet) values() []Rule {
	var ret []Rule
	for _, rule := range rs.rules {
		ret = append(ret, rule)
	}
	return ret
}

func (rs *ruleSet) Add(rule Rule) {
	rs.rules[rule.Name()] = rule
}

// RuleSet creates a new ruleset
func RuleSet(rules ...Rule) Rules {
	rs := new(ruleSet)
	rs.rules = make(map[RuleName]Rule)
	for _, rule := range rules {
		rs.rules[rule.Name()] = rule
	}
	return rs
}

// DefaultRule is a basic rule
type DefaultRule struct {
	name        RuleName
	description string
	condition   Condition
	actions     []Action
}

// Name gives the name of rule
func (r DefaultRule) Name() RuleName {
	return r.name
}

// Description gives the description of rule
func (r DefaultRule) Description() string {
	return r.description
}

func (r DefaultRule) evaluate(facts Facts) bool {
	return r.condition(facts)
}

func (r DefaultRule) execute(facts Facts) {
	execute(r.actions, facts)
}

func execute(actions []Action, facts Facts) {
	if len(actions) == 0 {
		return
	}
	action := actions[0]
	action(facts)
	execute(actions[1:], facts)
}

type defaultRuleBuilder struct {
	DefaultRule
}

func (b *defaultRuleBuilder) Given(condition Condition) RequiredActionOngoing {
	b.condition = condition
	return b
}

func (b *defaultRuleBuilder) Will(action Action) AdditionalActionOngoing {
	b.actions = []Action{action}
	return b
}

func (b *defaultRuleBuilder) Then(action Action) AdditionalActionOngoing {
	b.actions = append(b.actions, action)
	return b
}

func (b *defaultRuleBuilder) WithName(name string) DescriptionOngoing {
	b.name = RuleName(name)
	return b
}

func (b *defaultRuleBuilder) WithDescription(desc string) Outgoing {
	b.description = desc
	return b
}

func (b *defaultRuleBuilder) Build() Rule {
	defer func() {
		b = nil
	}()

	r := DefaultRule{}

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

	if zero.IsZeroVal(b.condition) {
		panic("rule condition is missing")
	} else {
		r.condition = b.condition
	}

	if zero.IsZeroVal(b.actions) || len(b.actions) == 0 {
		panic("rule actions is missing")
	} else {
		r.actions = b.actions
	}

	return r
}
