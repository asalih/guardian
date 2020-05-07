package operators

import "github.com/asalih/guardian/helpers"

//DataFileCaches Global *.data files caching variable
var DataFileCaches map[string]*DataFileCache

//DataFileCache *.data files caching model
type DataFileCache struct {
	FileName string
	Lines    []string
	Matcher  *helpers.Matcher
}
