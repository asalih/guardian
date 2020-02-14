package operators

func (opMap *OperatorMap) loadInspectFile() {
	opMap.funcMap["inspectFile"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
