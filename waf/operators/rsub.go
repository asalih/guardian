package operators

func (opMap *OperatorMap) loadRsub() {
	opMap.funcMap["rsub"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
