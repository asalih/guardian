package operators

import (
	GSQLI "github.com/koangel/grapeSQLI"
)

func (opMap *OperatorMap) loadDetectXSS() {
	opMap.funcMap["detectXSS"] = func(expression interface{}, variableData interface{}) bool {

		return GSQLI.XSSParser(variableData.(string))
	}
}
