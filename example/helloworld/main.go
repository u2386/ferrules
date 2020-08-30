package main

import (
	"fmt"

	"github.com/u2386/ferrules/pkg/ferrules"
)

var rules ferrules.Rules = ferrules.RuleSet(ferrules.
	Given(func(facts ferrules.Facts) bool {
		return true
	}).
	Will(func(facts ferrules.Facts) {
		fmt.Println("Hello World")
	}).
	Priority(1).
	WithName("Hello World rule").
	WithDescription("Always say hello world").
	Build(),
)

func main() {
	facts := ferrules.FactSet()
	engine := ferrules.DefaultRuleEngine{}
	engine.Fire(rules, facts)
}
