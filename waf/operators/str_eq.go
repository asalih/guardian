package operators

func init() {
	OperatorMaps.funcMap["streq"] = func(expression interface{}, variableData interface{}) bool {
		return expression.(string) == variableData.(string)
	}
}
