package livereload

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeJS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeJS)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "application/javascript"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expected)
	}

	// Check the response body is what we expect.
	expectedBody := liveReloadJS()
	if !bytes.Equal(rr.Body.Bytes(), expectedBody) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expectedBody))
	}
}
