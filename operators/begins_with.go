package operators

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadBeginsWith() {
	opMap.funcMap["beginsWith"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		return matches.NewMatchResult(false).SetMatch(strings.HasPrefix(variableData.(string), expression.(string)))
	}
}
