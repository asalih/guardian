package operators

func init() {
	OperatorMaps.funcMap["unconditionalMatch"] = func(expression interface{}, variableData interface{}) bool {
		return true
	}
}
