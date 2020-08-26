package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
)

// Rule is an interface that can be fired by engine
type Rule interface {
	evaluate(Facts) bool
	execute(Facts)
	Priority() RulePriority
	Name() RuleName
	Description() string
}

// Rules encapsulates a set of rules
type Rules interface {
	values() []Rule
	Add(Rule)
}

// Engine is an interface of rule engine
type Engine interface {
	Fire(Rules, Facts)
}
