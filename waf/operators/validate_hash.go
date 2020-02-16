package operators

func (opMap *OperatorMap) loadValidateHash() {
	opMap.funcMap["validateHash"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
