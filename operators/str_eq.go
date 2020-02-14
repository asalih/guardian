package operators

func (opMap *OperatorMap) loadStreq() {
	opMap.funcMap["streq"] = func(expression interface{}, variableData interface{}) bool {
		return expression.(string) == variableData.(string)
	}
}
