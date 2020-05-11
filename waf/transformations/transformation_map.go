package transformations

//TransformationMaps transformation map variable
var TransformationMaps *TransformationMap = &TransformationMap{make(map[string]func(interface{}) interface{})}

//TransformationMap struct mapper type
type TransformationMap struct {
	funcMap map[string]func(variableData interface{}) interface{}
}

//Get Reads the transformation func with given key
func (transform *TransformationMap) Get(key string) func(interface{}) interface{} {
	return transform.funcMap[key]
}
