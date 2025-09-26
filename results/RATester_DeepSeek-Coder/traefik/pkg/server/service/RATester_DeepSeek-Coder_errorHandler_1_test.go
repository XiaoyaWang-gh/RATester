package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestErrorHandler_1(t *testing.T) {
	t.Run("Test errorHandler", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		err := errors.New("test error")

		errorHandler(w, req, err)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != computeStatusCode(err) {
			t.Errorf("Expected status code %d, got %d", computeStatusCode(err), resp.StatusCode)
		}

		if string(body) != statusText(computeStatusCode(err)) {
			t.Errorf("Expected body %s, got %s", statusText(computeStatusCode(err)), string(body))
		}
	})
}
