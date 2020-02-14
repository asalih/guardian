package operators

func (opMap *OperatorMap) loadGeolookup() {
	opMap.funcMap["geolookup"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
