package engine

import (
	"github.com/asalih/guardian/matches"
)

func (t *TransactionMap) loadMultipartCrlfLfLines() *TransactionMap {
	t.variableMap["MULTIPART_CRLF_LF_LINES"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
