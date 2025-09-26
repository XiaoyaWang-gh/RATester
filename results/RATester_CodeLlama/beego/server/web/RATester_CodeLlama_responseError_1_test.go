package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	responseError(rw, r, 404, "404 page not found")
	if rw.Code != 404 {
		t.Errorf("responseError() code = %v, want %v", rw.Code, 404)
	}
	if rw.Body.String() != "404 page not found" {
		t.Errorf("responseError() body = %v, want %v", rw.Body.String(), "404 page not found")
	}
}
