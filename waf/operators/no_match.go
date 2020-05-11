package operators

func init() {
	OperatorMaps.funcMap["noMatch"] = func(expression interface{}, variableData interface{}) bool {
		return false
	}
}
