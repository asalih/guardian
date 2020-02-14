package models

import (
	"fmt"

	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/operators"
)

//SecRule VARIABLES OPERATOR [ACTIONS]

//Rule the rule model
type Rule struct {
	Variables []*Variable
	Operator  *Operator
	Action    *Action
	Chain     *Rule
}

//RuleExecutionResult the result object
type RuleExecutionResult struct {
	MatchResult *matches.MatchResult
	Rule        *Rule
}

//NewRule Inits a rule
func NewRule(variables []*Variable, operators *Operator, action *Action, chain *Rule) *Rule {
	return &Rule{variables, operators, action, chain}
}

//ExecuteRule Executes rule and returns match result
func (rule *Rule) ExecuteRule(variableData interface{}) *matches.MatchResult {
	matchResult := matches.NewMatchResult()

	fn := operators.OperatorMaps.Get(rule.Operator.Func)

	if fn == nil {
		//TODO Handle unrecognized fn
		fmt.Println("Unrecognized Operator fn" + rule.Operator.Func)
		return matches.NewMatchResult()
	}

	operatorResult := fn(rule.Operator.Expression, variableData)

	if operatorResult && !rule.Operator.OperatorIsNotType {
		return matchResult.SetMatch(true)
	} else if rule.Operator.OperatorIsNotType {
		return matchResult.SetMatch(true)
	}

	return matchResult
}

//ShouldBlock Determines whether rule is blocking action
func (rule *Rule) ShouldBlock() bool {
	return rule.Action.DisruptiveAction == DisruptiveActionBlock || rule.Action.DisruptiveAction == DisruptiveActionDeny
}
