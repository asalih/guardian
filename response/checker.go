package response

import (
	"net/http"

	"github.com/asalih/guardian/models"
)

//Checker Response checker
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	responseTransfer *responseTransfer
	result           chan *models.MatchResult
}

type responseTransfer struct {
	BodyBuffer  []byte
	ContentType string
}
