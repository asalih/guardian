package operators

func (opMap *OperatorMap) loadVerifyCC() {
	opMap.funcMap["verifyCC"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
