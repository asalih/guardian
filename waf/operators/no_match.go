package operators

func (opMap *OperatorMap) loadNoMatch() {
	opMap.funcMap["noMatch"] = func(expression interface{}, variableData interface{}) bool {
		return false
	}
}
