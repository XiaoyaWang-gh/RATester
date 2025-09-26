package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncodeQuerySemicolons_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "a=1;b=2" {
			t.Errorf("Expected query to be 'a=1;b=2', got '%s'", r.URL.RawQuery)
		}
	})

	req, _ := http.NewRequest("GET", "http://example.com/?a=1;b=2", nil)
	rr := httptest.NewRecorder()

	encodeQuerySemicolons(handler).ServeHTTP(rr, req)
}
