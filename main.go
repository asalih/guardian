package main

import (
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/waf/engine"
	"github.com/asalih/guardian/waf/operators"
	"github.com/asalih/guardian/waf/parser"
	"github.com/asalih/guardian/waf/transformations"
)

func init() {
	models.InitConfig()

	engine.InitTransactionMap()
	operators.InitOperatorMap()
	transformations.InitTransformationMap()

	parser.InitDataFiles()
	parser.InitRulesCollection()
}

func main() {

	srv := NewHTTPServer()
	srv.ServeHTTP()
}
