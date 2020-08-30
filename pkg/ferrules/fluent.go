package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
	"github.com/vikyd/zero"
)

// RequiredActionOngoing adds an action to a rule in a programming way
type RequiredActionOngoing interface {
	Will(Action) ActionOngoing
}

// ActionOngoing adds an action to a rule in a programming way
type ActionOngoing interface {
	Then(Action) ActionOngoing
	PriorityOngoing
}

// PriorityOngoing sets priority of rule
type PriorityOngoing interface {
	Priority(int) Outgoing
}

// Outgoing adds same descriptive attributes for a rule
type Outgoing interface {
	WithName(string) Outgoing
	WithDescription(string) Outgoing
	Build() Rule
}

type defaultRuleBuilder struct {
	DefaultRule
}

func (b *defaultRuleBuilder) Will(action Action) ActionOngoing {
	b.actions = []Action{action}
	return b
}

func (b *defaultRuleBuilder) Then(action Action) ActionOngoing {
	b.actions = append(b.actions, action)
	return b
}

func (b *defaultRuleBuilder) Priority(n int) Outgoing {
	b.priority = RulePriority(n)
	return b
}

func (b *defaultRuleBuilder) WithName(name string) Outgoing {
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

	if zero.IsZeroVal(b.priority) {
		panic("rule priority is missing")
	} else {
		r.priority = b.priority
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

// Given is the entry of fluent api
func Given(condition Condition) RequiredActionOngoing {
	b := &defaultRuleBuilder{}
	b.condition = condition
	return b
}
