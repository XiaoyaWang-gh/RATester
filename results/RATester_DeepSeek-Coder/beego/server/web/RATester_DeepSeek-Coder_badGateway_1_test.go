package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBadGateway_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/badGateway", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(badGateway)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadGateway)
	}

	expected := "<br>The page you have requested is down right now." +
		"<br><br><ul>" +
		"<br>The server, while acting as a gateway or proxy, received an invalid response from the upstream server it accessed in attempting to fulfill the request." +
		"<br>Please try again later and report the error to the website administrator" +
		"<br></ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
