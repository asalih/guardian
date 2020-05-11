package operators

import (
	GSQLI "github.com/koangel/grapeSQLI"
)

func init() {
	OperatorMaps.funcMap["detectSQLi"] = func(expression interface{}, variableData interface{}) bool {
		if err := GSQLI.SQLInject(variableData.(string)); err != nil {
			return true
		}

		return false
	}
}
