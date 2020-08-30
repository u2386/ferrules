package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
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
	priority    RulePriority
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

// Priority gives the priority of rule
func (r DefaultRule) Priority() RulePriority {
	return r.priority
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
