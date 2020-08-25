package main

import (
	"fmt"

	ferrules "github.com/u2386/ferrules/pkg"
)

func main() {
	rule := ferrules.NewRule("Hello World rule", "Always say hello world", 1, func(facts ferrules.Facts) bool {
		return true
	}, []ferrules.Action{func(facts ferrules.Facts) {
		fmt.Println("hello world")
	}})

	facts := make(map[string]ferrules.Fact)
	rules := map[ferrules.RuleName]ferrules.Rule{
		rule.Name(): rule,
	}
	engine := ferrules.DefaultRuleEngine{}
	ferrules.Run(&engine, rules, facts)
}
