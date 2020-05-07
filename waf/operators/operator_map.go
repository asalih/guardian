package operators

//RulesAndDatasPath static rules path
var RulesAndDatasPath string = "./crs/"

//OperatorMaps Global OperatorFuncs
var OperatorMaps *OperatorMap

//OperatorMap Map fn handler struct
type OperatorMap struct {
	funcMap map[string]func(expression interface{}, variableData interface{}) bool
}

//Get returns the operator fn with given key
func (ops *OperatorMap) Get(key string) func(interface{}, interface{}) bool {
	return ops.funcMap[key]
}

//InitOperatorMap operator initator
func InitOperatorMap() {
	OperatorMaps = &OperatorMap{make(map[string]func(interface{}, interface{}) bool)}

	OperatorMaps.loadBeginsWith()
	OperatorMaps.loadContains()
	OperatorMaps.loadContainsWord()
	OperatorMaps.loadDetectSqli()
	OperatorMaps.loadDetectXSS()
	OperatorMaps.loadEndsWith()
	OperatorMaps.loadEq()
	OperatorMaps.loadFuzzyHash() //Not implemented
	OperatorMaps.loadGe()
	OperatorMaps.loadGeolookup() //Not implemented
	OperatorMaps.loadGt()
	OperatorMaps.loadInspectFile() //Not implemented
	OperatorMaps.loadIPMatch()
	OperatorMaps.loadIPMatchFromFile() //Not implemented
	OperatorMaps.loadLe()
	OperatorMaps.loadLt()
	OperatorMaps.loadNoMatch()
	OperatorMaps.loadPm()
	OperatorMaps.loadPmFromFile()
	OperatorMaps.loadRbl()  //Not implemented
	OperatorMaps.loadRsub() //Not implemented
	OperatorMaps.loadRx()
	OperatorMaps.loadStreq()
	OperatorMaps.loadStrmatch()
	OperatorMaps.loadUnconditionalMatch()
	OperatorMaps.loadValidateByteRange()
	OperatorMaps.loadValidateDTD()    //Not implemented
	OperatorMaps.loadValidateHash()   //Not implemented
	OperatorMaps.loadValidateSchema() //Not implemented
	OperatorMaps.loadURLEncoding()
	OperatorMaps.loadValidateUtf8Encoding()
	OperatorMaps.loadVerifyCC()
	OperatorMaps.loadVerifyCPF()
	OperatorMaps.loadVerifySSN()
	OperatorMaps.loadWithin()
}
