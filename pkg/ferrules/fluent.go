package ferrules

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

// Incoming is the entry of fluent api
type Incoming interface {
	Given(condition Condition) RequiredActionOngoing
}
