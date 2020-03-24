package bodyprocessor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antchfx/xmlquery"
)

//XMLBodyProcessor URL Encoded body parser
type XMLBodyProcessor struct {
	request      *http.Request
	bodyBuffer   []byte
	hasBodyError bool
	XMLDocument  *xmlquery.Node
}

//GetBody ...
func (p *XMLBodyProcessor) GetBody() map[string][]string {

	if p.XMLDocument != nil {
		return nil
	}

	doc, err := xmlquery.Parse(bytes.NewBuffer(p.GetBodyBuffer()))

	if err != nil {
		fmt.Println(err)

		p.hasBodyError = true
	}

	p.XMLDocument = doc

	return nil
}

//GetBodyBuffer ...
func (p *XMLBodyProcessor) GetBodyBuffer() []byte {

	if len(p.bodyBuffer) > 0 {
		return p.bodyBuffer
	}

	bodyBytes, _ := ioutil.ReadAll(p.request.Body)
	p.request.Body.Close() //  must close
	p.bodyBuffer = bodyBytes
	p.request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return p.bodyBuffer
}

//HasBodyError ...
func (p *XMLBodyProcessor) HasBodyError() bool {
	return p.hasBodyError
}
