package operators

import (
	"strconv"

	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadRx() {
	opMap.funcMap["rx"] = func(expression interface{}, variableData interface{}) bool {

		switch variableData.(type) {
		case string:
			isMatch, _ := helpers.IsMatch(expression.(string), variableData.(string))
			return isMatch
		case int:
			isMatch, _ := helpers.IsMatch(expression.(string), strconv.Itoa(variableData.(int)))
			return isMatch
		}

		return false
	}
}
