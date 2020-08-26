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
	Then(func(facts ferrules.Facts) {
		if fact, err := facts.Get("name"); err == nil {
			fmt.Printf("Hello, %v\n", fact.Value.(string))
		}
	}).
	Then(func(facts ferrules.Facts) {
		if fact, err := facts.Get("name"); err == nil {
			fact.Value = "Hugo"
			facts.Add(fact)
		}
	}).
	Then(func(facts ferrules.Facts) {
		if fact, err := facts.Get("name"); err == nil {
			fmt.Printf("Hello, %v\n", fact.Value.(string))
		}
	}).
	Priority(1).
	WithName("Hello World rule").
	WithDescription("Always say hello world").
	Build(),
)

func main() {
	facts := ferrules.FactSet()
	facts.Put("name", "Leon")
	engine := ferrules.DefaultRuleEngine{}
	engine.Fire(rules, facts)
}
