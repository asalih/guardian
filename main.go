package main

import (
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/waf/parser"
	"github.com/asalih/guardian/waf/transformations"
)

func init() {
	models.InitConfig()

	transformations.InitTransformationMap()

	parser.InitDataFiles()
	parser.InitRulesCollection()
}

func main() {

	srv := NewHTTPServer()
	srv.ServeHTTP()
}
