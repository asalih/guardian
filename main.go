package main

import (
	"Guardian/Models"
)

func main() {

	//Let's init the payload data collection
	models.InitPayloadDataCollection()

	srv := &HTTPServer{}

	srv.ServeHTTP()
}
