package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiProxyByType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{
		// Initialize your service here
	}

	req, err := http.NewRequest("GET", "/api/proxy/type/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(svr.apiProxyByType)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"proxies":[{"name":"proxy1","type":"test","status":"active"},{"name":"proxy2","type":"test","status":"active"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
