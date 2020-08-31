package types

// RuleName represents the name of a rule
type RuleName string

func (name RuleName) String() string {
	return string(name)
}

// RulePriority represents the priority of a rule
type RulePriority int

// Int converts priority to int
func (p RulePriority) Int() int {
	return int(p)
}

// FactName represents the name of a fact
type FactName string

func (name FactName) String() string {
	return string(name)
}
