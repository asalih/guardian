package operators

import (
	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadRx() {
	opMap.funcMap["rx"] = func(expression interface{}, variableData interface{}) bool {
		isMatch, _ := helpers.IsMatch(expression.(string), variableData.(string))

		return isMatch
	}
}
