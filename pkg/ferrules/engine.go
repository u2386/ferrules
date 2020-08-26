package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
)

// DefaultRuleEngine iterates over a set of rules which
// sorted by their priority in nature order, evaluates the condition
// of each rule, executes actions if the condition met
type DefaultRuleEngine struct {
}

// Fire triggers all rules on given facts
func (engine *DefaultRuleEngine) Fire(rules Rules, facts Facts) {
	engine.fire(values(rules), facts)
}

func (engine *DefaultRuleEngine) fire(rules []Rule, facts Facts) {
	if len(rules) == 0 {
		return
	}
	rule := rules[0]
	if ok := rule.evaluate(facts); ok {
		rule.execute(facts)
	}
	engine.fire(rules[1:], facts)
}

func values(rules map[RuleName]Rule) []Rule {
	var ret []Rule
	for _, rule := range rules {
		ret = append(ret, rule)
	}
	return ret
}
