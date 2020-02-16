package main

import (
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/waf/engine"
	"github.com/asalih/guardian/waf/operators"
	"github.com/asalih/guardian/waf/parser"
	"github.com/asalih/guardian/waf/transformations"
)

func main() {

	models.InitConfig()

	engine.InitTransactionMap()
	operators.InitOperatorMap()
	transformations.InitTransformationMap()

	parser.InitRulesCollection()

	srv := NewHTTPServer()
	srv.ServeHTTP()
}
