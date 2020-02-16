package operators

import (
	"github.com/koangel/grapeSQLI"
)

func (opMap *OperatorMap) loadDetectXss() {
	opMap.funcMap["detectXSS"] = func(expression interface{}, variableData interface{}) bool {

		return GSQLI.XSSParser(variableData.(string))
	}
}
