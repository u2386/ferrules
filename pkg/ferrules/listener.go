package ferrules

// RuleEngineListener is a listener for engine execution events
type RuleEngineListener struct {
	Before func(Rules, Facts)
	After  func(Rules, Facts)
}
