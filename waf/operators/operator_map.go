package operators

//RulesAndDatasPath static rules path
var RulesAndDatasPath string = "./crs/"

//OperatorMaps Global OperatorFuncs
var OperatorMaps *OperatorMap = &OperatorMap{make(map[string]func(interface{}, interface{}) bool)}

//OperatorMap Map fn handler struct
type OperatorMap struct {
	funcMap map[string]func(expression interface{}, variableData interface{}) bool
}

//Get returns the operator fn with given key
func (ops *OperatorMap) Get(key string) func(interface{}, interface{}) bool {
	return ops.funcMap[key]
}
