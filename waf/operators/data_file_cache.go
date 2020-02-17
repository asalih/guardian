package operators

var DataFileCaches map[string]*DataFileCache

type DataFileCache struct {
	FileName string
	Lines    []string
}
