package operators

func init() {
	OperatorMaps.funcMap["validateSchema"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
