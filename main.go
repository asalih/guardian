package main

import "github.com/asalih/guardian/models"

func main() {
	models.InitConfig()

	//Let's init the payload data collection
	models.InitRequestPayloadDataCollection()
	models.InitResponsePayloadDataCollection()

	srv := NewHTTPServer()

	srv.ServeHTTP()
}
