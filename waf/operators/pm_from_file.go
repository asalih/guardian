package operators

import (
	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadPmFromFile() {
	fn := func(expression interface{}, variableData interface{}) bool {

		fileCache := DataFileCaches[expression.(string)]

		if fileCache == nil {
			return false
		}

		m := helpers.NewStringMatcher(fileCache.Lines)
		hits := m.Match([]byte(variableData.(string)))

		if len(hits) > 0 {
			return true
		}

		return false
	}

	opMap.funcMap["pmf"] = fn
	opMap.funcMap["pmFromFile"] = fn
}
