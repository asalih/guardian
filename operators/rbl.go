package operators

func (opMap *OperatorMap) loadRbl() {
	opMap.funcMap["rbl"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
