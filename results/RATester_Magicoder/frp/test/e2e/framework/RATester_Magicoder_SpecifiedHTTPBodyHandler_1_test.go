package framework

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSpecifiedHTTPBodyHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testBody := []byte("test body")
	handler := SpecifiedHTTPBodyHandler(testBody)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := testBody
	if !bytes.Equal(rr.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}
