package transformations

//TransformationMaps transformation map variable
var TransformationMaps *TransformationMap

//TransformationMap struct mapper type
type TransformationMap struct {
	funcMap map[string]func(variableData interface{}) interface{}
}

//Get Reads the transformation func with given key
func (transform *TransformationMap) Get(key string) func(interface{}) interface{} {
	return transform.funcMap[key]
}

//InitTransformationMap inits transformation maps
func InitTransformationMap() {
	TransformationMaps = &TransformationMap{make(map[string]func(interface{}) interface{})}

	TransformationMaps.loadRemoveWhitespace()
	TransformationMaps.loadBase64Decode()
	TransformationMaps.loadSQLHexDecode()    // Not implemented
	TransformationMaps.loadBase64DecodeExt() // Not implemented
	TransformationMaps.loadBase64Encode()
	TransformationMaps.loadCmdLine() // Not implemented
	TransformationMaps.loadCompressWhitespace()
	TransformationMaps.loadEscapeSeqDecode() // Not implemented
	TransformationMaps.loadHexDecode()
	TransformationMaps.loadHexEncode()
	TransformationMaps.loadHTMLEntityDecode()
	TransformationMaps.loadJSDecode() // Not implemented
	TransformationMaps.loadLength()
	TransformationMaps.loadLowercase()
	TransformationMaps.loadMD5()
	TransformationMaps.loadNone()
	TransformationMaps.loadNormalizePath() // Not implemented
	TransformationMaps.loadParityEven7bit()
	TransformationMaps.loadParityOdd7bit()
	TransformationMaps.loadParityZero7bit()
	TransformationMaps.loadRemoveNulls()
	TransformationMaps.loadReplaceComments()
	TransformationMaps.loadRemoveCommentsChar()
	TransformationMaps.loadReplaceNulls()
	TransformationMaps.loadURLDecode()
	TransformationMaps.loadURLDecodeUni() //TODO Review it
	TransformationMaps.loadUppercase()
	TransformationMaps.loadUtf8ToUnicode() // Not implemented
	TransformationMaps.loadSHA1()
	TransformationMaps.loadTrimLeft()
	TransformationMaps.loadTrimRight()
}
