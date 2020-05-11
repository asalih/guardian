package operators

func init() {
	OperatorMaps.funcMap["validateDTD"] = func(expression interface{}, variableData interface{}) bool {
		//TODO: might have to review
		return false
	}
}
