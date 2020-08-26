package ferrules

// Condition represents a rule's condition
type Condition func(facts Facts) bool
