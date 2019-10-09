package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

//Config Application settings
type Config struct {
	ConnectionString string `json:"connectionString"`
}

//Configuration ...
var Configuration Config

//InitConfig initializes the config file
func InitConfig() {
	cnfFile := "appsettings"

	guardianEnv := os.Getenv("GUARDIAN_ENV")

	if guardianEnv != "" {
		cnfFile += "." + strings.ToLower(guardianEnv)
	} else {
		cnfFile += ".development"
	}

	cnfFile += ".json"

	jsonFile, err := ioutil.ReadFile(cnfFile)

	jerr := json.Unmarshal(jsonFile, &Configuration)

	if err != nil || jerr != nil {
		panic("Configuration file error. " + cnfFile)
	}
}
