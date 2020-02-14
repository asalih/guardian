package operators

func (opMap *OperatorMap) loadValidateSchema() {
	opMap.funcMap["validateSchema"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
