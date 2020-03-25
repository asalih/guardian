package engine

import (
	"github.com/antchfx/xmlquery"
	"github.com/asalih/guardian/matches"
	"github.com/asalih/guardian/waf/bodyprocessor"
)

func (t *TransactionMap) loadXML() *TransactionMap {
	t.variableMap["XML"] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {

			switch executer.transaction.RequestBodyProcessor.(type) {
			case *bodyprocessor.XMLBodyProcessor:
				executer.transaction.RequestBodyProcessor.GetBody()
				bodyProcessor := executer.transaction.RequestBodyProcessor.(*bodyprocessor.XMLBodyProcessor)

				if bodyProcessor.HasBodyError() {
					return matches.NewMatchResult()
				}

				var nodes []*xmlquery.Node
				if len(executer.variable.Filter) > 0 {
					nodes = xmlquery.Find(bodyProcessor.XMLDocument, executer.variable.Filter[0])
				} else {
					nodes = []*xmlquery.Node{bodyProcessor.XMLDocument}
				}
				for _, node := range nodes {
					match := executer.rule.ExecuteRule(node.InnerText())

					if match.IsMatched {
						return match
					}
				}
			}

			return matches.NewMatchResult()

		}}

	return t
}
