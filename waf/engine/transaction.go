package engine

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asalih/guardian/matches"

	"github.com/asalih/guardian/models"
)

var TransactionMaps *TransactionMap

type TransactionMap struct {
	variableMap map[string]*TransactionData
}

type Transaction struct {
	Request  *http.Request
	Response *http.Response
	tx       map[string]interface{}
}

type TransactionData struct {
	executer func(*TransactionExecuterModel) *matches.MatchResult
}

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
	TransactionMaps.loadRequestCookies()
	TransactionMaps.loadRequestCookiesNames()
	TransactionMaps.loadRequestFilename()
	TransactionMaps.loadRequestHeaders()
	TransactionMaps.loadRequestHeadersNames()
	TransactionMaps.loadRequestBodyType()
	TransactionMaps.loadRequestUri()
	TransactionMaps.loadRequestMethod()
	TransactionMaps.loadRemoteAddr()
	TransactionMaps.loadTX()
	TransactionMaps.loadIP()  //Not implemented
	TransactionMaps.loadXML() //Not implemented
}

// NewTransaction Initiates a new request variable object
func NewTransaction(r *http.Request) *Transaction {
	return &Transaction{r, nil, make(map[string]interface{})}
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

//SafeParseForm ReadAll functions clears the buffer while parsing the form body.
//For preventing to loose body buffer we are caching it first.
func (t *Transaction) SafeParseForm() error {
	if t.Request.Form != nil && t.Request.PostForm != nil {
		return nil
	}

	bodyBytes, _ := ioutil.ReadAll(t.Request.Body)
	t.Request.Body.Close() //  must close
	t.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	err := t.Request.ParseForm()
	t.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return err
}
