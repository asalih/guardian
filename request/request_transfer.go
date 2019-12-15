package request

import (
	"mime"
	"net/http"

	"github.com/asalih/guardian/models"
	"github.com/asalih/guardian/variables"
)

//RequestTransfer Request variable object
type RequestTransfer struct {
	varMap map[string]interface{}
}

// NewRequestTransfer Initiates a new request variable object
func NewRequestTransfer(request *http.Request) *RequestTransfer {
	transferData := RequestTransfer{make(map[string]interface{})}

	transferData.varMap[variables.REQUEST_COOKIES] = request.Cookies()
	transferData.varMap[variables.REQUEST_URI] = request.RequestURI

	isHTTPS := request.TLS != nil

	transferData.varMap[variables.REQUEST_URI] = request.RequestURI

	if isHTTPS {
		transferData.varMap[variables.REQUEST_URI_RAW] = "https://" + request.Host + request.RequestURI
		transferData.varMap[variables.REQUEST_PROTOCOL] = "https"
	} else {
		transferData.varMap[variables.REQUEST_URI_RAW] = "http://" + request.Host + request.RequestURI
		transferData.varMap[variables.REQUEST_PROTOCOL] = "http"
	}

	transferData.varMap[variables.QUERY_STRING] = models.UnEscapeRawValue(request.URL.RawQuery)
	transferData.varMap[variables.REQUEST_BASENAME] = request.URL.Path
	transferData.varMap[variables.REQUEST_COOKIES_NAMES] = models.GetCookiesNames(request.Cookies())
	transferData.varMap[variables.REQUEST_HEADERS] = request.Header
	transferData.varMap[variables.REQUEST_HEADERS_NAMES] = models.GetHeadersNames(request.Header)

	contentType := request.Header.Get("Content-Type")

	mediaType, mediaParams, err := mime.ParseMediaType(contentType)
	transferData.varMap[variables.REQUEST_BODY_TYPE] = mediaType
	transferData.varMap[variables.MULTIPART_BOUNDARY] = mediaParams["boundary"]

	if err != nil && contentType != "" {
		transferData.varMap[variables.MULTIPART_ERROR] = 1
	} else {
		transferData.varMap[variables.MULTIPART_ERROR] = 0
	}

	return &transferData
}

// Get Returns a value with given key
func (transfer *RequestTransfer) Get(key string) string {
	return transfer.varMap[key].(string)
}

// GetMap Returns a value map with given key
func (transfer *RequestTransfer) GetMap(key string) map[string]string {
	return transfer.varMap[key].(map[string]string)
}
