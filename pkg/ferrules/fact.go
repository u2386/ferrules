package ferrules

import (
	"fmt"

	. "github.com/u2386/ferrules/internal/types"
)

// Fact is a mutable reality truth
type Fact struct {
	Name  FactName
	Value interface{}
}

// FactSet creates a new facts set
func FactSet() Facts {
	return &factSet{
		facts: make(map[FactName]*Fact),
	}
}

type factSet struct {
	facts map[FactName]*Fact
}

func (fs *factSet) Put(name string, value interface{}) {
	fs.Add(&Fact{
		Name:  FactName(name),
		Value: value,
	})
}

func (fs *factSet) Add(fact *Fact) {
	fs.facts[fact.Name] = fact
}

func (fs *factSet) Remove(name string) {
	for key := range fs.facts {
		if string(key) == name {
			delete(fs.facts, key)
		}
	}
}

func (fs *factSet) Get(name string) (*Fact, error) {
	for key := range fs.facts {
		if string(key) == name {
			f := fs.facts[key]
			return f, nil
		}
	}
	return nil, fmt.Errorf("%s not found", name)
}
