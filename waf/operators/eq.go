package operators

import (
	"strconv"
)

func (opMap *OperatorMap) loadEq() {
	opMap.funcMap["eq"] = func(expression interface{}, variableData interface{}) bool {

		expInt, errExpr := strconv.Atoi(expression.(string))
		varInt := 0
		var errVar error

		switch v := variableData.(type) {
		case string:
			varInt, errVar = strconv.Atoi(v)
		case int:
			varInt = v
		}

		if errExpr != nil || errVar != nil {
			return false
		}

		return expInt == varInt
	}
}
