package transformations

var TransformationMaps *TransformationMap

type TransformationMap struct {
	funcMap map[string]func(variableData interface{}) interface{}
}

func (transform *TransformationMap) Get(key string) func(interface{}) interface{} {
	return transform.funcMap[key]
}

func InitTransformationMap() {
	TransformationMaps = &TransformationMap{make(map[string]func(interface{}) interface{})}

	TransformationMaps.loadLowercase()
	TransformationMaps.loadRemoveWhitespace()
	TransformationMaps.loadUrlDecode()
	TransformationMaps.loadUrlDecodeUni() //TODO Review it
	TransformationMaps.loadNone()
	TransformationMaps.loadBase64Decode()
	TransformationMaps.loadSQLHexDecode()    // Not implemented
	TransformationMaps.loadBase64DecodeExt() // Not implemented
	TransformationMaps.loadBase64Encode()
}
