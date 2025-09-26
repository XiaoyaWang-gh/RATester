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

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	notImplemented(rw, r)
	if rw.Code != 501 {
		t.Errorf("Expected status code 501, got %v", rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is Not Implemented."+
		"<br><br><ul>"+
		"<br>Please try again later and report the error to the website administrator"+
		"<br></ul>" {
		t.Errorf("Expected body %v, got %v", "<br>The page you have requested is Not Implemented."+
			"<br><br><ul>"+
			"<br>Please try again later and report the error to the website administrator"+
			"<br></ul>", rw.Body.String())
	}
}
