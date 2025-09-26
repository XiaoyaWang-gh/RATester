package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDenyFragment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	testHandler := denyFragment(handler)

	req, err := http.NewRequest("GET", "http://example.com/path#fragment", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
