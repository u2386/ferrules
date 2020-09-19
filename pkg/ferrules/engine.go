package ferrules

import (
	"github.com/vikyd/zero"
)

// DefaultRuleEngine iterates over a set of rules which
// sorted by their priority in nature order, evaluates the condition
// of each rule, executes actions if the condition met
type DefaultRuleEngine struct {
	Listener []RuleEngineListener
}

// Fire triggers all rules on given facts
func (engine *DefaultRuleEngine) Fire(rules Rules, facts Facts) {
	for _, listener := range engine.Listener {
		if !zero.IsZeroVal(listener.Before) {
			listener.Before(rules, facts)
		}
	}

	fire(engine, rules.values(), facts)

	for _, listener := range engine.Listener {
		if !zero.IsZeroVal(listener.After) {
			listener.After(rules, facts)
		}
	}
}

func fire(engine Engine, rules []Rule, facts Facts) {
	if len(rules) == 0 {
		return
	}
	rule := rules[0]
	if ok := rule.evaluate(facts); ok {
		rule.execute(facts)
	}
	fire(engine, rules[1:], facts)
}
