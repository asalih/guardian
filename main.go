package main

import (
	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/waf/parser"
)

func init() {
	models.InitConfig()

	parser.InitDataFiles()
	parser.InitRulesCollection()
}

func main() {

	srv := NewHTTPServer()
	srv.ServeHTTP()
}
