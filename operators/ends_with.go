package operators

import (
	"strings"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadEndsWith() {
	opMap.funcMap["endsWith"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		matchResult := matches.NewMatchResult(false)

		return matchResult.SetMatch(strings.HasSuffix(variableData.(string), expression.(string)))
	}
}
