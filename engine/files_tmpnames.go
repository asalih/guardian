package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES_TMPNAMES = "FILES_TMPNAMES"

func (t *TransactionMap) loadFilesTmpNames() *TransactionMap {
	t.variableMap[FILES_TMPNAMES] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			//TODO Not implemented yet
			return matches.NewMatchResult()
		}}

	return t
}
