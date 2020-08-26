package ferrules

import (
	. "github.com/u2386/ferrules/internal/types"
)

// Fact is a mutable reality truth
type Fact struct {
	Name  FactName
	Value interface{}
}

// Facts encapsulates a set of facts
type Facts map[string]Fact
