package operators

func init() {
	OperatorMaps.funcMap["inspectFile"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		//NA yet!
		return false
	}
}
