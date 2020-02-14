package engine

import (
	"fmt"
	"net/http"

	"github.com/asalih/guardian/matches"

	"github.com/asalih/guardian/models"
)

var TransactionMaps *TransactionMap

type TransactionMap struct {
	variableMap map[string]*TransactionData
}

type Transaction struct {
	request *http.Request
}

type TransactionData struct {
	executer func(*TransactionExecuterModel) *matches.MatchResult
}

type TransactionExecuterModel struct {
	request         *http.Request
	transactionData *TransactionData
	rule            *models.Rule
	variable        *models.Variable
}

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
	TransactionMaps.loadRequestCookies()
	TransactionMaps.loadRequestCookiesNames()
	TransactionMaps.loadRequestHeaders()
	TransactionMaps.loadRequestHeadersNames()
	TransactionMaps.loadRequestBodyType()
	TransactionMaps.loadRequestUri()
}

// NewRequestTransfer Initiates a new request variable object
func NewTransaction(request *http.Request) *Transaction {
	return &Transaction{request}
}

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

		executerModel := &TransactionExecuterModel{t.request, mapData, rule, variable}
		matchResult = mapData.executer(executerModel)

		if matchResult.IsMatched && !variable.FilterIsNotType && !rule.Operator.OperatorIsNotType {
			if rule.Chain != nil {
				matchResult = t.Execute(rule.Chain)

				if matchResult == nil {
					continue
				}

				if matchResult.IsMatched {
					return matchResult
				}
			} else {
				return matchResult
			}
		} else if !matchResult.IsMatched && !matchResult.DefaultState && (variable.FilterIsNotType || rule.Operator.OperatorIsNotType) {
			return matchResult.SetMatch(true)
		}
	}

	return matchResult

}
