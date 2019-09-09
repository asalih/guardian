package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//PayloadDataCollection parsed json file
var PayloadDataCollection []PayloadData

//PayloadData for checking requests
type PayloadData struct {
	CheckPoint string `json:"checkPoint"`
	Payload    string `json:"payload"`
	Type       string `json:"type"`
}

//InitPayloadDataCollection Payload data initializer
func InitPayloadDataCollection() {
	jsonFile, err := ioutil.ReadFile("requestPayloads.json")

	if err != nil {
		panic(err)
	}

	jerr := json.Unmarshal(jsonFile, &PayloadDataCollection)

	if jerr != nil {
		fmt.Println(jerr)
		panic(jerr)
	}
}
