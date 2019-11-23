package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//RequestPayloadDataCollection parsed json file
var RequestPayloadDataCollection []PayloadData

//ResponsePayloadDataCollection parsed json file
var ResponsePayloadDataCollection []PayloadData

//RequestCheckPointPayloadData grouped by checkpoint
var RequestCheckPointPayloadData map[string][]PayloadData = make(map[string][]PayloadData)

//ResponseCheckPointPayloadData grouped by checkpoint
var ResponseCheckPointPayloadData map[string][]PayloadData = make(map[string][]PayloadData)

//LenOfGroupedRequestPayloadDataCollection parsed payload data length
var LenOfGroupedRequestPayloadDataCollection int

//LenOfGroupedResponsePayloadDataCollection parsed payload data length
var LenOfGroupedResponsePayloadDataCollection int

//PayloadData for checking requests
type PayloadData struct {
	Action      string `json:"action"`
	CheckPoint  string `json:"checkPoint"`
	Payload     string `json:"payload"`
	Type        string `json:"type"`
	MatchResult *PayloadData
}

//InitRequestPayloadDataCollection Payload data initializer
func InitRequestPayloadDataCollection() {
	InitRequestPayloadDataCollectionFile("requestPayloads.json")
}

//InitRequestPayloadDataCollectionFile Payload data initializer
func InitRequestPayloadDataCollectionFile(path string) {
	jsonFile, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	jerr := json.Unmarshal(jsonFile, &RequestPayloadDataCollection)

	if jerr != nil {
		fmt.Println(jerr)
		panic(jerr)
	}

	for _, m := range RequestPayloadDataCollection {
		RequestCheckPointPayloadData[m.CheckPoint] = append(RequestCheckPointPayloadData[m.CheckPoint], m)
	}

	LenOfGroupedRequestPayloadDataCollection = len(RequestCheckPointPayloadData)
}

//InitResponsePayloadDataCollectionFile Payload data initializer
func InitResponsePayloadDataCollection() {
	InitResponsePayloadDataCollectionFile("responsePayloads.json")
}

//InitResponsePayloadDataCollectionFile Payload data initializer
func InitResponsePayloadDataCollectionFile(path string) {
	jsonFile, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	jerr := json.Unmarshal(jsonFile, &ResponsePayloadDataCollection)

	if jerr != nil {
		fmt.Println(jerr)
		panic(jerr)
	}

	for _, m := range ResponsePayloadDataCollection {
		ResponseCheckPointPayloadData[m.CheckPoint] = append(ResponseCheckPointPayloadData[m.CheckPoint], m)
	}

	LenOfGroupedResponsePayloadDataCollection = len(ResponseCheckPointPayloadData)
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
