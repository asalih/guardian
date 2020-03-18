package engine

import (
	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/waf/bodyprocessor"
)

func (t *TransactionMap) loadReqBodyProcessor() *TransactionMap {
	t.variableMap["REQBODY_PROCESSOR"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			var body string

			switch executer.transaction.BodyProcessor.(type) {
			case *bodyprocessor.JSONBodyProcessor:
				body = "JSON"
				break
			case *bodyprocessor.MultipartProcessor:
				body = "MULTIPART"
				break
			case *bodyprocessor.URLEncodedProcessor:
				body = "URLENCODED"
				break
			case *bodyprocessor.XMLBodyProcessor:
				body = "XML"
				break
			}

			return executer.rule.ExecuteRule(body)
		}}

	return t
}
