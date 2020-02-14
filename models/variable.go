package models

import (
	"github.com/asalih/guardian/helpers"
)

//Variable definition for a rule
type Variable struct {
	Name                     string
	Filter                   []string
	FilterIsNotType          bool
	LengthCheckForCollection bool
}

//ShouldPassCheck Variable filter
func (variable *Variable) ShouldPassCheck(value string) bool {
	if variable.Filter == nil {
		return false
	}

	filterContainsKey := helpers.StringContains(variable.Filter, value)

	if !filterContainsKey && !variable.FilterIsNotType ||
		filterContainsKey && variable.FilterIsNotType {
		return true
	}

	return false
}
