package operators

func (opMap *OperatorMap) loadFuzzyHash() {
	opMap.funcMap["fuzzyHash"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
