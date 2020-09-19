package main

import (
	"fmt"
	"github.com/u2386/ferrules/pkg/ferrules"
)

var rules ferrules.Rules = ferrules.RuleSet(
	ferrules.CompositeRuleBuilder().
		AnyOf(
			ferrules.RuleBuilder().
				Given(func(facts ferrules.Facts) bool {
					fact, _ := facts.Get("value")
					return fact.Value.(int)%3 == 0
				}).
				Will(func(facts ferrules.Facts) {
					facts.Put("echo", "fizz")
				}).
				Priority(1).
				WithName("Fizz Rule").
				WithDescription("echo `fizz` if value is a multiple of 3").
				Build(),

			ferrules.RuleBuilder().
				Given(func(facts ferrules.Facts) bool {
					fact, _ := facts.Get("value")
					return fact.Value.(int)%5 == 0
				}).
				Will(func(facts ferrules.Facts) {
					facts.Put("echo", "buzz")
				}).
				Priority(1).
				WithName("Buzz Rule").
				WithDescription("echo `buzz` if value is a multiple of 5").
				Build(),

			ferrules.RuleBuilder().
				Given(func(facts ferrules.Facts) bool {
					return true
				}).
				Will(func(facts ferrules.Facts) {
					fact, _ := facts.Get("value")
					facts.Put("echo", fact.Value)
				}).
				Priority(1).
				WithName("Default Rule").
				WithDescription("echo value").
				Build(),
		).
		Priority(1).
		WithName("FizzBuzz Game").
		WithDescription("fizzbuzz puzzle game").
		Build(),
)

func main() {
	engine := ferrules.DefaultRuleEngine{}

	for i := 1; i < 10; i++ {
		facts := ferrules.FactSet()
		facts.Put("value", i)
		engine.Fire(rules, facts)
		fact, _ := facts.Get("echo")
		fmt.Println(fact.Value)
	}
}
