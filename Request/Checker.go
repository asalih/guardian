package request

import (
	models "Guardian/Models"
	"fmt"
	"net/http"
	"sync"
)

/*Checker Cheks the requests init*/
type Checker struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	result chan bool
}

/*NewRequestChecker Request checker initializer*/
func NewRequestChecker(w http.ResponseWriter, r *http.Request) *Checker {

	return &Checker{w, r, nil}
}

/*Handle Request checker handler func*/
func (r Checker) Handle() bool {

	done := make(chan bool, 1)

	go func() {
		var wg sync.WaitGroup
		lengthOfPayloads := len(models.PayloadDataCollection)
		r.result = make(chan bool, lengthOfPayloads)

		wg.Add(lengthOfPayloads)

		for _, payload := range models.PayloadDataCollection {
			go r.handlePayload(payload, &wg)
		}

		wg.Wait()

		close(r.result)

		done <- true
	}()

	<-done

	for i := range r.result {
		fmt.Println(i)
		if i {
			r.ResponseWriter.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(r.ResponseWriter, "Bad Request. %s", r.Request.URL.Path)

			return true
		}
	}
	fmt.Println("exiting")
	return false
}

func (r Checker) handlePayload(payload models.PayloadData, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(payload.Payload)
	isMatch, _ := models.IsMatch(payload.Payload, models.UnEscapeRawValue(r.Request.URL.RawQuery))

	r.result <- isMatch
}
