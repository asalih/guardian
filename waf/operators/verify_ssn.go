package operators

func (opMap *OperatorMap) loadVerifySSN() {
	opMap.funcMap["verifySSN"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
