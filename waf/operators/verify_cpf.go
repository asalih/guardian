package operators

func (opMap *OperatorMap) loadVerifyCPF() {
	opMap.funcMap["verifyCPF"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
