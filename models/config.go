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
	RateLimitSec     int    `json:"rateLimitSec"`
	RateLimitBurst   int    `json:"rateLimitBurst"`
}

//Configuration ...
var Configuration Config

//InitConfig ...
func InitConfig() {
	InitConfigFile("appsettings")
}

//InitConfigFile initializes the config file
func InitConfigFile(cnfFile string) {

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
		panic(err)
	}
}
