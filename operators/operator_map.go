package operators

import (
	"github.com/asalih/guardian/matches"
)

var OperatorMaps *OperatorMap

type OperatorMap struct {
	funcMap map[string]func(expression interface{}, variableData interface{}) *matches.MatchResult
}

func (ops *OperatorMap) Get(key string) func(interface{}, interface{}) *matches.MatchResult {
	return ops.funcMap[key]
}

func InitOperatorMap() {
	OperatorMaps = &OperatorMap{make(map[string]func(interface{}, interface{}) *matches.MatchResult)}

	OperatorMaps.loadBeginsWith()
	OperatorMaps.loadContains()
	OperatorMaps.loadContainsWord()
	OperatorMaps.loadDetectSqli()
	OperatorMaps.loadDetectXss()
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
	OperatorMaps.loadPmFromFile() //Not implemented
	OperatorMaps.loadRbl()        //Not implemented
	OperatorMaps.loadRsub()       //Not implemented
	OperatorMaps.loadRx()
	OperatorMaps.loadStreq()
	OperatorMaps.loadStrmatch()
	OperatorMaps.loadUnconditionalMatch()
	OperatorMaps.loadValidateByteRange()
	OperatorMaps.loadValidateDTD()    //Not implemented
	OperatorMaps.loadValidateHash()   //Not implemented
	OperatorMaps.loadValidateSchema() //Not implemented
	OperatorMaps.loadUrlEncoding()
	OperatorMaps.loadValidateUtf8Encoding()
	OperatorMaps.loadVerifyCC()
	OperatorMaps.loadVerifyCPF()
	OperatorMaps.loadVerifySSN()
	OperatorMaps.loadWithin()
}
