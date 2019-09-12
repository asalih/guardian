package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//PayloadDataCollection parsed json file
var PayloadDataCollection []PayloadData

//CheckPointPayloadData grouped by checkpoint
var CheckPointPayloadData map[string][]PayloadData = make(map[string][]PayloadData)

//LenOfGroupedPayloadDataCollection parsed payload data length
var LenOfGroupedPayloadDataCollection int

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

	//CheckPointPayloadData := make(map[string][]PayloadData)

	for _, m := range PayloadDataCollection {
		CheckPointPayloadData[m.CheckPoint] = append(CheckPointPayloadData[m.CheckPoint], m)
	}

	LenOfGroupedPayloadDataCollection = len(CheckPointPayloadData)
}

//Filter filters the payload data
func Filter(ss []PayloadData, test func(PayloadData) bool) (ret []PayloadData) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
