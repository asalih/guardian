package operators

import (
	GSQLI "github.com/koangel/grapeSQLI"
)

func init() {
	OperatorMaps.funcMap["detectXSS"] = func(expression interface{}, variableData interface{}) bool {

		return GSQLI.XSSParser(variableData.(string))
	}
}
