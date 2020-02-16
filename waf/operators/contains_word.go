package operators

import (
	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadContainsWord() {
	opMap.funcMap["containsWord"] = func(expression interface{}, variableData interface{}) bool {
		isMatch, _ := helpers.IsMatch("\b"+expression.(string)+"\b", variableData.(string))

		return isMatch
	}
}
