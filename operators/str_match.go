package operators

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadStrmatch() {
	opMap.funcMap["strmatch"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(strings.Contains(variableData.(string), expression.(string)))
	}
}
