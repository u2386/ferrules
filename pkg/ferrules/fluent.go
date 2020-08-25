package ferrules

// RequiredActionOngoing adds an action to a rule in a programming way
type RequiredActionOngoing interface {
	Will(Action) ActionOngoing
}

// ActionOngoing adds an action to a rule in a programming way
type ActionOngoing interface {
	Will(Action) ActionOngoing
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
	rule
}

func (b *builder) Will(action Action) ActionOngoing {
	b.rule.actions = append(b.rule.actions, action)
	return b
}

func (b *builder) Priority(n int) Outgoing {
	b.rule.priority = RulePriority(n)
	return b
}

func (b *builder) WithName(name string) Outgoing {
	b.rule.name = RuleName(name)
	return b
}

func (b *builder) WithDescription(desc string) Outgoing {
	b.rule.description = desc
	return b
}

func (b *builder) Build() Rule {
	r := b.rule
	return &r
}

// Given is the entry of fluent api
func Given(condition Condition) RequiredActionOngoing {
	return &builder{
		rule: rule{
			name:        "",
			description: "",
			condition:   condition,
			actions:     []Action{},
			priority:    0,
		},
	}
}
