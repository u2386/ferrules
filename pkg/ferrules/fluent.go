package ferrules

// RequiredActionOngoing adds an action to a rule in a programming way
type RequiredActionOngoing interface {
	Will(Action) AdditionalActionOngoing
}

// AdditionalActionOngoing adds an action to a rule in a programming way
type AdditionalActionOngoing interface {
	Then(Action) AdditionalActionOngoing
	NameOngoing
}

// NameOngoing adds name to a rule
type NameOngoing interface {
	WithName(string) DescriptionOngoing
}

// DescriptionOngoing sets a description for a rule
type DescriptionOngoing interface {
	WithDescription(string) Outgoing
}

// Outgoing is the terminal of fluent api
type Outgoing interface {
	Build() Rule
}

// Incoming is the entry of fluent api
type Incoming interface {
	Given(condition Condition) RequiredActionOngoing
}

// CompositeIncoming is the entry of fluent composite api
type CompositeIncoming interface {
	AnyOf(rules ...Rule) NameOngoing
}
