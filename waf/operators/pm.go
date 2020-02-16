package operators

import (
	"strings"

	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadPm() {
	opMap.funcMap["pm"] = func(expression interface{}, variableData interface{}) bool {

		keywords := strings.Split(expression.(string), " ")
		m := helpers.NewStringMatcher(keywords)
		hits := m.Match([]byte(variableData.(string)))

		if len(hits) > 0 {
			return true
		}

		return false
	}
}
