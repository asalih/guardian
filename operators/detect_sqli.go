package operators

import (
	"github.com/koangel/grapeSQLI"
)

func (opMap *OperatorMap) loadDetectSqli() {
	opMap.funcMap["detectSQLi"] = func(expression interface{}, variableData interface{}) bool {
		if err := GSQLI.SQLInject(variableData.(string)); err != nil {
			return true
		}

		return false
	}
}
