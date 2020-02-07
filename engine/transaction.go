package engine

import (
	"net/http"

	"github.com/asalih/guardian/matches"

	"github.com/asalih/guardian/models"
)

type Transaction struct {
	variableMap map[string]*TransactionData
	request     *http.Request
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

// NewRequestTransfer Initiates a new request variable object
func NewTransaction(request *http.Request) *Transaction {
	transactionData := Transaction{make(map[string]*TransactionData), request}

	transactionData.loadArgs()
	transactionData.loadArgsNames()
	transactionData.loadArgsCombinedSize()
	transactionData.loadAuthType()
	transactionData.loadDuration() //Not implemented
	transactionData.loadEnv()      //Not implemented
	transactionData.loadFiles()
	transactionData.loadFilesCombinedSize() //TODO might add mime type
	transactionData.loadFilesNames()
	transactionData.loadFullRequestAndLength()
	transactionData.loadFilesSizes()
	transactionData.loadFilesTmpNames()        //Not implemented
	transactionData.loadFilesTmpContent()      //Not implemented
	transactionData.loadGeo()                  //Not implemented
	transactionData.loadHighestSeverity()      //Not implemented
	transactionData.loadInboundDataError()     //Not implemented
	transactionData.loadMatchedVar()           //Not implemented
	transactionData.loadMatchedVars()          //Not implemented
	transactionData.loadMultipartCrlfLfLines() //Not implemented
	transactionData.loadMultipartName()
	transactionData.loadQueryString()
	transactionData.loadRequestCookies()
	transactionData.loadRequestCookiesNames()
	transactionData.loadRequestHeaders()
	transactionData.loadRequestHeadersNames()
	transactionData.loadRequestBodyType()
	transactionData.loadRequestUri()

	return &transactionData
}

func (t *Transaction) ExecuteRule(rule *models.Rule) *matches.MatchResult {

	var matchResult *matches.MatchResult

	for _, variable := range rule.Variables {
		mapData := t.variableMap[variable.Name]

		if mapData == nil {
			//TODO log unknown Rule
			return matches.NewMatchResult(false)
		}

		executerModel := &TransactionExecuterModel{t.request, mapData, rule, &variable}
		matchResult = mapData.executer(executerModel)

		if matchResult.IsMatched {
			return matchResult
		}
	}

	return matchResult

}
