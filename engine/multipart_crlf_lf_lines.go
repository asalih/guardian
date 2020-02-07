package engine

import (
	"github.com/asalih/guardian/matches"
)

var MULTIPART_CRLF_LF_LINES = "MULTIPART_CRLF_LF_LINES"

func (t *Transaction) loadMultipartCrlfLfLines() *Transaction {
	t.variableMap[MULTIPART_CRLF_LF_LINES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult(false)
		}}

	return t
}
