package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
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

type builder struct {
	R *rule
}

func (b *builder) Will(action Action) ActionOngoing {
	b.R.actions = append(b.R.actions, action)
	return b
}

func (b *builder) Then(action Action) ActionOngoing {
	b.R.actions = append(b.R.actions, action)
	return b
}

func (b *builder) Priority(n int) Outgoing {
	b.R.priority = RulePriority(n)
	return b
}

func (b *builder) WithName(name string) Outgoing {
	b.R.name = RuleName(name)
	return b
}

func (b *builder) WithDescription(desc string) Outgoing {
	b.R.description = desc
	return b
}

func (b *builder) Build() Rule {
	r := b.R
	b.R = nil
	return r
}

// Given is the entry of fluent api
func Given(condition Condition) RequiredActionOngoing {
	return &builder{
		R: &rule{
			name:        "",
			description: "",
			condition:   condition,
			actions:     []Action{},
			priority:    0,
		},
	}
}
