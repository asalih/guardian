package models

import (
	"github.com/asalih/guardian/helpers"
	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/operators"
)

//SecRule VARIABLES OPERATOR [ACTIONS]

//Rule the rule model
type Rule struct {
	Variables []Variable
	Operators []Operator
	Actions   []string
}

//Operator definition for a rule
type Operator struct {
	Func              string
	Expression        string
	OperatorIsNotType bool
}

//Variable definition for a rule
type Variable struct {
	Name                     string
	Filter                   []string
	FilterIsNotType          bool
	LengthCheckForCollection bool
}

func (variable *Variable) ShouldPassCheck(value string) bool {
	filterContainsKey := helpers.StringContains(variable.Filter, value)

	if !filterContainsKey && !variable.FilterIsNotType ||
		filterContainsKey && variable.FilterIsNotType {
		return true
	}

	return false
}

//NewRule Inits a rule
func NewRule(variables []Variable, operators []Operator, actions string) *Rule {
	return &Rule{variables, operators, []string{actions}}
}

func (rule *Rule) ExecuteRule(variableData interface{}) *matches.MatchResult {

	var matchResult *matches.MatchResult

	for _, ops := range rule.Operators {
		fn := operators.OperatorMaps.Get(ops.Func)

		if fn == nil {
			//TODO Handle unrecognized fn
			return matches.NewMatchResult(false)
		}

		matchResult = fn(ops.Expression, variableData)

		if matchResult.IsMatched && !ops.OperatorIsNotType {
			return matchResult
		} else if ops.OperatorIsNotType {
			return matchResult.SetMatch(true)
		}
	}

	return matchResult
}
