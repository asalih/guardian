package operators

import (
	"bufio"
	"os"
	"strings"

	"github.com/asalih/guardian/helpers"
)

func (opMap *OperatorMap) loadPmFromFile() {
	fn := func(expression interface{}, variableData interface{}) bool {

		dataFile, err := os.Open(RulesAndDatasPath + expression.(string))

		if err != nil {
			return false
		}

		keywords := make([]string, 0)
		scanner := bufio.NewScanner(dataFile)
		for scanner.Scan() {
			word := scanner.Text()

			if word == "" || strings.HasPrefix(word, "#") {
				continue
			}

			keywords = append(keywords, word)
		}

		m := helpers.NewStringMatcher(keywords)
		hits := m.Match([]byte(variableData.(string)))

		if len(hits) > 0 {
			return true
		}

		return false
	}

	opMap.funcMap["pmf"] = fn
	opMap.funcMap["pmFromFile"] = fn
}
