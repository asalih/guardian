package main

import "github.com/asalih/guardian/models"

func main() {

	//Let's init the payload data collection
	models.InitRequestPayloadDataCollection()
	models.InitResponsePayloadDataCollection()

	srv := &HTTPServer{}

	srv.ServeHTTP()
}
