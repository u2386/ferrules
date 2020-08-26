package main

import (
	"fmt"

	. "github.com/u2386/ferrules/internal/types"

	"github.com/u2386/ferrules/pkg/ferrules"
)

func main() {
	rule := ferrules.
		Given(func(facts ferrules.Facts) bool {
			return true
		}).
		Will(func(facts ferrules.Facts) {
			fmt.Println("hello world")
		}).
		Then(func(facts ferrules.Facts) {
			fmt.Println("hello leon")
		}).
		Priority(1).
		WithName("Hello World rule").
		WithDescription("Always say hello world").
		Build()

	facts := make(map[string]ferrules.Fact)
	rules := map[RuleName]ferrules.Rule{
		rule.Name(): rule,
	}
	engine := ferrules.DefaultRuleEngine{}
	engine.Fire(rules, facts)
}
