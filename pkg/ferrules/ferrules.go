package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
)

// Rule is an interface that can be fired by engine
type Rule interface {
	evaluate(Facts) bool
	execute(Facts)
	Name() RuleName
	Description() string
}

// Rules provides an interface for a set of rules
type Rules interface {
	values() []Rule
	Add(Rule)
}

// Facts provides an interface for a set of facts
type Facts interface {
	Put(string, interface{})
	Add(*Fact)
	Remove(string)
	Get(string) (*Fact, error)
}

// Engine is an interface of rule engine
type Engine interface {
	Fire(Rules, Facts)
}

// RuleBuilder create a new rule builder
func RuleBuilder() Incoming {
	return &defaultRuleBuilder{}
}

// CompositeRuleBuilder create a composite rule builder
func CompositeRuleBuilder() CompositeIncoming {
	return &compositeRuleBuilder{}
}
