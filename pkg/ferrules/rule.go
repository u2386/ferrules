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

type rule struct {
	name        RuleName
	description string
	priority    RulePriority
	condition   Condition
	actions     []Action
}

func (r *rule) Name() RuleName {
	return r.name
}

func (r *rule) Description() string {
	return r.description
}

func (r *rule) Priority() RulePriority {
	return r.priority
}

func (r *rule) evaluate(facts Facts) bool {
	return r.condition(facts)
}

func (r *rule) execute(facts Facts) {
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
