package parser

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/asalih/guardian/helpers"

	"github.com/asalih/guardian/waf/operators"
)

//InitDataFiles Data files initor
func InitDataFiles() {
	operators.DataFileCaches = make(map[string]*operators.DataFileCache)

	files, _ := ioutil.ReadDir(operators.RulesAndDatasPath)
	for _, v := range files {
		if v.IsDir() || !strings.HasSuffix(v.Name(), ".data") {
			continue
		}

		initDataFile(v.Name())
	}
}

//InitRulesCollectionFile Rules data initializer
func initDataFile(name string) {
	dataFile, err := os.Open(operators.RulesAndDatasPath + name)

	if err != nil {
		panic(err)
	}

	fileCache := &operators.DataFileCache{FileName: name}
	scanner := bufio.NewScanner(dataFile)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 1 || strings.HasPrefix(line, "#") {
			continue
		}

		readLine := strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(line), "\r"), "\n", " ")

		if len(readLine) <= 1 {
			continue
		}

		fileCache.Lines = append(fileCache.Lines, readLine)
	}

	if len(fileCache.Lines) > 0 {
		fileCache.Matcher = helpers.NewStringMatcher(fileCache.Lines)
	}
	operators.DataFileCaches[name] = fileCache
}
