package operators

func (opMap *OperatorMap) loadValidateDTD() {
	opMap.funcMap["validateDTD"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
