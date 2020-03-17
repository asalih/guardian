package bodyprocessor

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/lloyd/goj"
)

//JSONBodyProcessor JSON body parser
type JSONBodyProcessor struct {
	request    *http.Request
	bodyBuffer []byte
	cache      map[string][]string
}

//GetBody ...
func (p *JSONBodyProcessor) GetBody() map[string][]string {
	if len(p.cache) > 0 {
		return p.cache
	}

	parser := goj.NewParser()

	arrIdx := 0
	inArr := false
	path := []string{}

	gFunc := goj.Callback(func(what goj.Type, key, value []byte) goj.Action {
		keyStr := string(key)
		switch what {
		case goj.String, goj.False, goj.True, goj.Float, goj.Integer, goj.NegInteger, goj.Null:
			cp := strings.Join(path, ".")
			if inArr {
				cp = cp + ".array_" + strconv.Itoa(arrIdx)
				arrIdx++
			}

			if cp != "" {
				if keyStr != "" {
					cp = cp + "." + keyStr
				}
			} else {
				cp = keyStr
			}

			p.cache[cp] = []string{string(value)}

		case goj.Array:
			inArr = true
			if keyStr != "" {
				path = append(path, keyStr)
			}
		case goj.Object:
			if keyStr != "" {
				path = append(path, keyStr)
			}
		case goj.ArrayEnd:
			inArr = false
			arrIdx = 0
			if len(path) > 0 {
				path = path[:len(path)-1]
			}
		case goj.ObjectEnd:
			if len(path) > 0 {
				path = path[:len(path)-1]
			}
		}

		return goj.Continue
	})

	err := parser.Parse(p.GetBodyBuffer(), gFunc)

	if err != nil {
		panic(err)
	}

	return p.cache
}

//GetBodyBuffer ...
func (p *JSONBodyProcessor) GetBodyBuffer() []byte {

	if len(p.bodyBuffer) > 0 {
		return p.bodyBuffer
	}

	bodyBytes, _ := ioutil.ReadAll(p.request.Body)
	p.request.Body.Close() //  must close
	p.bodyBuffer = bodyBytes

	return p.bodyBuffer
}
