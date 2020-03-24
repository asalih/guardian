package engine

import (
	"fmt"
	"net/http"

	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/waf/bodyprocessor"

	"github.com/asalih/guardian/models"
)

//TransactionMaps global map variable object
var TransactionMaps *TransactionMap

//TransactionMap variable map model
type TransactionMap struct {
	variableMap map[string]*TransactionData
}

//Transaction Main model for request response examine
type Transaction struct {
	Request               *http.Request
	Response              *http.Response
	RequestBodyProcessor  bodyprocessor.IBodyProcessor
	ResponseBodyProcessor bodyprocessor.IBodyProcessor

	tx map[string]interface{}
}

//TransactionData Transaction model
type TransactionData struct {
	executer func(*TransactionExecuterModel) *matches.MatchResult
}

//TransactionExecuterModel Executer model
type TransactionExecuterModel struct {
	transaction *Transaction
	rule        *models.Rule
	variable    *models.Variable
}

//InitTransactionMap inits transaction map fns
func InitTransactionMap() {
	TransactionMaps = &TransactionMap{make(map[string]*TransactionData)}

	TransactionMaps.loadArgs()
	TransactionMaps.loadArgsNames()
	TransactionMaps.loadArgsCombinedSize()
	TransactionMaps.loadAuthType()
	TransactionMaps.loadDuration() //Not implemented
	TransactionMaps.loadEnv()      //Not implemented
	TransactionMaps.loadFiles()
	TransactionMaps.loadFilesCombinedSize() //TODO might add mime type
	TransactionMaps.loadFilesNames()
	TransactionMaps.loadFullRequestAndLength()
	TransactionMaps.loadFilesSizes()
	TransactionMaps.loadFilesTmpNames()        //Not implemented
	TransactionMaps.loadFilesTmpContent()      //Not implemented
	TransactionMaps.loadGeo()                  //Not implemented
	TransactionMaps.loadHighestSeverity()      //Not implemented
	TransactionMaps.loadInboundDataError()     //Not implemented
	TransactionMaps.loadMatchedVar()           //Not implemented
	TransactionMaps.loadMatchedVars()          //Not implemented
	TransactionMaps.loadMultipartCrlfLfLines() //Not implemented
	TransactionMaps.loadQueryString()
	TransactionMaps.loadUniqueID() // Not implemented - might not needed.
	TransactionMaps.loadRequestBody()
	TransactionMaps.loadReqBodyError()
	TransactionMaps.loadRequestBodyLength()
	TransactionMaps.loadRequestCookies()
	TransactionMaps.loadRequestCookiesNames()
	TransactionMaps.loadRequestFilename()
	TransactionMaps.loadRequestHeaders()
	TransactionMaps.loadRequestHeadersNames()
	TransactionMaps.loadRequestBodyType()
	TransactionMaps.loadReqBodyProcessor()
	TransactionMaps.loadRequestLine()
	TransactionMaps.loadRequestURI()
	TransactionMaps.loadRequestMethod()
	TransactionMaps.loadRemoteAddr()
	TransactionMaps.loadTX()
	TransactionMaps.loadIP() //Not implemented
	TransactionMaps.loadXML()
	TransactionMaps.loadResponseStatus()
	TransactionMaps.loadResponseBody()
	TransactionMaps.loadResponseBodyLength()
}

// NewTransaction Initiates a new request variable object
func NewTransaction(r *http.Request) *Transaction {
	return &Transaction{r, nil, bodyprocessor.NewBodyProcessor(r), nil, make(map[string]interface{})}
}

//Get the data in transaction data
func (tMap *TransactionMap) Get(key string) *TransactionData {
	return tMap.variableMap[key]
}

//Execute Executes transaction for rule
func (t *Transaction) Execute(rule *models.Rule) *matches.MatchResult {

	var matchResult *matches.MatchResult

	for _, variable := range rule.Variables {
		mapData := TransactionMaps.Get(variable.Name)

		if mapData == nil {
			//TODO log unknown Rule
			fmt.Println("Unrecognized variable: " + variable.Name)
			return nil
		}

		executerModel := &TransactionExecuterModel{t, rule, variable}
		matchResult = mapData.executer(executerModel)

		if matchResult.IsMatched {
			if rule.Chain != nil {
				matchResult = t.Execute(rule.Chain)

				if matchResult == nil {
					continue
				}
			}

			if !variable.FilterIsNotType && !rule.Operator.OperatorIsNotType {
				return matchResult
			} else if !matchResult.DefaultState {
				matchResult.SetMatch(false)
			}

		} else if !matchResult.IsMatched && !matchResult.DefaultState && (variable.FilterIsNotType || rule.Operator.OperatorIsNotType) {
			return matchResult.SetMatch(true)
		}
	}

	return matchResult

}
