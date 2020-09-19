package main

import (
	"fmt"

	"github.com/u2386/ferrules/pkg/ferrules"
)

var rules = ferrules.RuleSet(ferrules.RuleBuilder().
	Given(func(facts ferrules.Facts) bool {
		rain, _ := facts.Get("rain")
		return rain.Value.(bool)
	}).
	Will(func(facts ferrules.Facts) {
		fmt.Println("It's rain, take an umbrella!")
	}).
	Then(func(facts ferrules.Facts) {
		facts.Put("umbrella", "open")
	}).
	WithName("weather rule").
	WithDescription("if it rains then take an umbrella").
	Build())

func main() {
	facts := ferrules.FactSet()
	facts.Put("rain", true)

	engine := ferrules.DefaultRuleEngine{}
	engine.Fire(rules, facts)
	fmt.Println(facts)
}
