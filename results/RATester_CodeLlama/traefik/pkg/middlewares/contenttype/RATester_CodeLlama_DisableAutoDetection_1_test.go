package contenttype

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDisableAutoDetection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if _, ok := rw.Header()["Content-Type"]; ok {
			t.Error("Content-Type should not be set")
		}
	})

	handler := DisableAutoDetection(next)

	rw := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handler.ServeHTTP(rw, req)
}
