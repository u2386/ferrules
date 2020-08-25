package ferrules

// FactName represents the name of a fact
type FactName string

// RuleName represents the name of a rule
type RuleName string

// RulePriority represents the priority of a rule
type RulePriority int

type rule struct {
	name        RuleName
	description string
	priority    RulePriority
	condition   Condition
	actions     []Action
}

// NewRule creates a rule
func NewRule(name string, desc string, priority int, condition Condition, actions []Action) Rule {
	return &rule{
		name:        RuleName(name),
		description: desc,
		priority:    RulePriority(priority),
		condition:   condition,
		actions:     actions,
	}
}

func (r *rule) Name() RuleName {
	return r.name
}

func (r *rule) Description() string {
	return r.description
}

func (r *rule) Priority() RulePriority {
	return r.priority
}

func (r *rule) evaluate(facts Facts) bool {
	return r.condition(facts)
}

func (r *rule) execute(facts Facts) {
	execute(r.actions, facts)
}

func execute(actions []Action, facts Facts) {
	if len(actions) == 0 {
		return
	}
	action := actions[0]
	action(facts)
	execute(actions[1:], facts)
}

// Fact is a mutable reality truth
type Fact struct {
	Name  FactName
	Value interface{}
}

// Facts encapsulates a set of facts
type Facts map[string]Fact

// Action represents a rule's action
type Action func(facts Facts)

// Condition represents a rule's condition
type Condition func(facts Facts) bool

// Rule is an interface that can be fired by engine
type Rule interface {
	evaluate(Facts) bool
	execute(Facts)
	Priority() RulePriority
	Name() RuleName
	Description() string
}

// Rules encapsulates a set of rules
type Rules map[RuleName]Rule

// Engine is an interface of rule engine
type Engine interface {
	fire([]Rule, Facts)
}

// DefaultRuleEngine iterates over a set of rules which
// sorted by their priority in nature order, evaluates the condition
// of each rule, executes actions if the condition met
type DefaultRuleEngine struct {
}

func (engine *DefaultRuleEngine) fire(rules []Rule, facts Facts) {
	if len(rules) == 0 {
		return
	}
	rule := rules[0]
	if ok := rule.evaluate(facts); ok {
		rule.execute(facts)
	}
	engine.fire(rules[1:], facts)
}

// Run is an API for triggering rule engine
func Run(engine Engine, rules Rules, facts Facts) {
	values := func(rules map[RuleName]Rule) []Rule {
		var ret []Rule
		for _, rule := range rules {
			ret = append(ret, rule)
		}
		return ret
	}
	engine.fire(values(rules), facts)
}
