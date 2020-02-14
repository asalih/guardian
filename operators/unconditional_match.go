package operators

func (opMap *OperatorMap) loadUnconditionalMatch() {
	opMap.funcMap["unconditionalMatch"] = func(expression interface{}, variableData interface{}) bool {
		return true
	}
}
