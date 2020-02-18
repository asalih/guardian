package operators

import "github.com/asalih/guardian/helpers"

var DataFileCaches map[string]*DataFileCache

type DataFileCache struct {
	FileName string
	Lines    []string
	Matcher  *helpers.Matcher
}
