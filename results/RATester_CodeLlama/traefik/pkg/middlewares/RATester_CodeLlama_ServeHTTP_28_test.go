package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_28(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &httptest.ResponseRecorder{}
	req := &http.Request{}
	h := &HTTPHandlerSwitcher{}
	h.ServeHTTP(rw, req)
}
