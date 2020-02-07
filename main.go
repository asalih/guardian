package main

import (

	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/operators"
)

func main() {

	models.InitConfig()

	operators.InitOperatorMap()
	models.InitRulesCollection()

	srv := NewHTTPServer()

	srv.ServeHTTP()
}
