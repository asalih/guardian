package operators

import (
	"github.com/asalih/guardian/helpers"
)

func init() {
	OperatorMaps.funcMap["containsWord"] = func(expression interface{}, variableData interface{}) bool {
		isMatch, _ := helpers.IsMatch("\b"+expression.(string)+"\b", variableData.(string))

		return isMatch
	}
}
