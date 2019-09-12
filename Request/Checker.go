package request

import (
	models "Guardian/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

var staticSuffix = []string{".js", ".css", ".png", ".jpg", ".gif", ".bmp", ".svg", ".ico"}

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	requestTransfer *requestTransfer
	result          chan bool
}

type requestTransfer struct {
	BodyBuffer  []byte
	ContentType string
}

/*NewRequestChecker Request checker initializer*/
func NewRequestChecker(w http.ResponseWriter, r *http.Request) *Checker {

	return &Checker{w, r, nil, nil}
}

/*Handle Request checker handler func*/
func (r Checker) Handle() bool {

	if r.Request.Method == "GET" && r.IsStaticResource(r.Request.URL.Path) {
		return false
	}

	bodyBuffer, _ := ioutil.ReadAll(r.Request.Body)
	r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))
	r.requestTransfer = &requestTransfer{bodyBuffer, r.Request.Header.Get("Content-Type")}

	done := make(chan bool, 1)

	go func() {
		var wg sync.WaitGroup

		r.result = make(chan bool, models.LenOfGroupedPayloadDataCollection)

		wg.Add(models.LenOfGroupedPayloadDataCollection)

		for key, payload := range models.CheckPointPayloadData {
			go r.handlePayload(key, payload, &wg)
		}

		wg.Wait()

		close(r.result)

		done <- true
	}()

	<-done

	for i := range r.result {
		if i {
			r.ResponseWriter.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

			return true
		}
	}

	return false
}

// IsStaticResource ...
func (r Checker) IsStaticResource(url string) bool {
	if strings.Contains(url, "?") {
		return false
	}
	for _, suffix := range staticSuffix {
		if strings.HasSuffix(url, suffix) {
			return true
		}
	}
	return false
}

func (r Checker) handlePayload(key string, payloads []models.PayloadData, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(r.Request.RequestURI)
	switch key {
	case "Query":
		r.result <- r.handleQuery(payloads)
	case "Path":
		r.result <- r.handlePath(payloads)
	case "Form":
		r.result <- r.handleForm(payloads)
	case "Upload":
		//Upload check point will be handled in handleForm function
		r.result <- false
	default:
		r.result <- false
	}
}

func (r Checker) handleQuery(payloads []models.PayloadData) bool {
	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(r.Request.URL.RawQuery))

		if isMatch {
			return true
		}
	}

	return false
}

func (r Checker) handlePath(payloads []models.PayloadData) bool {
	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, r.Request.URL.Path)

		if isMatch {
			return true
		}
	}

	return false
}

func (r Checker) handleForm(payloads []models.PayloadData) bool {
	mediaType, mediaParams, _ := mime.ParseMediaType(r.requestTransfer.ContentType)

	if strings.HasPrefix(mediaType, "multipart/form-data") {
		// ChkPoint_UploadFileExt
		r.Request.ParseMultipartForm(1024)
		if r.Request.MultipartForm != nil {
			for _, filesHeader := range r.Request.MultipartForm.File {
				for _, fileHeader := range filesHeader {
					fileExtension := filepath.Ext(fileHeader.Filename) // .php
					uploadCheck := models.Filter(models.PayloadDataCollection, func(p models.PayloadData) bool { return p.CheckPoint == "Upload" })

					for i := 0; i < len(uploadCheck); i++ {
						matched, _ := models.IsMatch(uploadCheck[i].Payload, fileExtension)
						if matched == true {
							return matched
						}
					}
				}
			}

			// Multipart Content
			body1 := ioutil.NopCloser(bytes.NewBuffer(r.requestTransfer.BodyBuffer))
			multiReader := multipart.NewReader(body1, mediaParams["boundary"])
			for {
				p, err := multiReader.NextPart()
				if err == io.EOF {
					break
				}
				partContent, err := ioutil.ReadAll(p)

				for _, p := range payloads {
					isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(string(partContent)))

					if isMatch {
						return true
					}
				}
			}
		}

	} else if strings.HasPrefix(mediaType, "application/json") {
		var params interface{}
		err := json.Unmarshal(r.requestTransfer.BodyBuffer, &params)

		if err != nil {
			panic(err)
		}
		//TODO: Handle json
		// matched, policy := IsJSONValueHitPolicy(ctxMap, appID, params)
		// if matched == true {
		// 	return matched, policy
		// }
	} else {
		r.Request.ParseForm()
	}

	params := r.Request.Form

	r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(r.requestTransfer.BodyBuffer))
	for key, values := range params {

		for _, p := range payloads {
			isMatch, _ := models.IsMatch(p.Payload, key)

			if isMatch {
				return true
			}
		}

		for _, value := range values {
			if isDigit, err := models.IsMatch(`^\d{1,5}$`, value); err == nil {
				if isDigit {
					continue
				}
			}
			// ChkPoint_ValueLength
			// TODO: Check length
			// valueLength := strconv.Itoa(len(value))
			// matched, policy = IsMatchGroupPolicy(ctxMap, appID, valueLength, models.ChkPointValueLength, "", false)
			// if matched == true {
			// 	return matched, policy
			// }
			// ChkPoint_GetPostValue

			for _, p := range payloads {
				isMatch, _ := models.IsMatch(p.Payload, value)

				if isMatch {
					return true
				}
			}

			return false
		}

	}

	return false
}
