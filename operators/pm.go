package operators

import (
	"strings"

	"github.com/asalih/guardian/helpers"

	"github.com/asalih/guardian/matches"
)

func (opMap *OperatorMap) loadPm() {
	opMap.funcMap["pm"] = func(expression interface{}, variableData interface{}) *matches.MatchResult {
		result := matches.NewMatchResult(false)

		keywords := strings.Split(expression.(string), " ")
		m := helpers.NewStringMatcher(keywords)
		hits := m.Match([]byte(variableData.(string)))

		if len(hits) > 0 {
			result.SetMatch(true)
		}

		return result
	}
}
