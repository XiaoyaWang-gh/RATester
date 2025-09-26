package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotImplemented_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/notImplemented", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(notImplemented)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotImplemented)
	}

	expected := "<br>The page you have requested is Not Implemented." +
		"<br><br><ul>" +
		"<br>Please try again later and report the error to the website administrator" +
		"<br></ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
