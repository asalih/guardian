package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"guardian/models"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"
)

var staticSuffix = []string{".js", ".css", ".png", ".jpg", ".gif", ".bmp", ".svg", ".ico"}

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	requestTransfer *requestTransfer
	result          chan *models.MatchResult
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

		r.result = make(chan *models.MatchResult, models.LenOfGroupedPayloadDataCollection)

		wg.Add(models.LenOfGroupedPayloadDataCollection)

		for key, payload := range models.CheckPointPayloadData {
			go r.handlePayload(key, payload, &wg)
		}

		wg.Wait()

		close(r.result)

		done <- true
	}()

	select {
	case res := <-done:
		fmt.Println(res)
	case <-time.After(3 * time.Minute):
		panic("failed to execute rules.")
	}

	for i := range r.result {
		if i.IsMatched {
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
		r.result <- models.NewMatchResult(nil, false)
	default:
		r.result <- models.NewMatchResult(nil, false)
	}
}

func (r Checker) handleQuery(payloads []models.PayloadData) *models.MatchResult {
	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(r.Request.URL.RawQuery))

		if isMatch {
			return models.NewMatchResult(&p, true)
		}
	}

	return models.NewMatchResult(nil, false)
}

func (r Checker) handlePath(payloads []models.PayloadData) *models.MatchResult {
	for _, p := range payloads {
		isMatch, _ := models.IsMatch(p.Payload, r.Request.URL.Path)

		if isMatch {
			return models.NewMatchResult(&p, true)
		}
	}

	return models.NewMatchResult(nil, false)
}

func (r Checker) handleForm(payloads []models.PayloadData) *models.MatchResult {
	mediaType, mediaParams, _ := mime.ParseMediaType(r.requestTransfer.ContentType)

	if strings.HasPrefix(mediaType, "multipart/form-data") {
		// ChkPoint_UploadFileExt
		r.Request.ParseMultipartForm(1024)
		if r.Request.MultipartForm != nil {
			for _, filesHeader := range r.Request.MultipartForm.File {
				for _, fileHeader := range filesHeader {
					fileExtension := filepath.Ext(fileHeader.Filename) // .php
					uploadCheck := models.Filter(models.PayloadDataCollection, func(p models.PayloadData) bool { return p.CheckPoint == "Upload" })

					for _, p := range uploadCheck {
						matched, _ := models.IsMatch(p.Payload, fileExtension)
						if matched == true {
							return models.NewMatchResult(&p, true)
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
						return models.NewMatchResult(&p, true)
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

		result := r.isJSONValueHitPolicy(payloads, params)
		if result.IsMatched {
			return result
		}
	} else {
		r.Request.ParseForm()
	}

	params := r.Request.Form

	r.Request.Body = ioutil.NopCloser(bytes.NewBuffer(r.requestTransfer.BodyBuffer))
	for key, values := range params {

		for _, p := range payloads {
			isMatch, _ := models.IsMatch(p.Payload, key)

			if isMatch {
				return models.NewMatchResult(&p, true)
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
					return models.NewMatchResult(&p, true)
				}
			}

			return models.NewMatchResult(nil, false)
		}

	}

	return models.NewMatchResult(nil, false)
}

func (r Checker) isJSONValueHitPolicy(payloads []models.PayloadData, value interface{}) *models.MatchResult {
	valueKind := reflect.TypeOf(value).Kind()
	switch valueKind {
	case reflect.String:
		value2 := value.(string)

		for _, p := range payloads {
			isMatch, _ := models.IsMatch(p.Payload, models.UnEscapeRawValue(value2))

			if isMatch {
				return models.NewMatchResult(&p, true)
			}
		}
	case reflect.Map:
		value2 := value.(map[string]interface{})
		for _, subValue := range value2 {
			matched := r.isJSONValueHitPolicy(payloads, subValue)
			if matched.IsMatched {
				return matched
			}
		}
	case reflect.Slice:
		value2 := value.([]interface{})
		for _, subValue := range value2 {
			result := r.isJSONValueHitPolicy(payloads, subValue)
			if result.IsMatched {
				return result
			}
		}
	}
	return models.NewMatchResult(nil, false)
}
