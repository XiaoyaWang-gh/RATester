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

	rw := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	badGateway(rw, r)
	if rw.Code != 502 {
		t.Errorf("expected status code 502, got %v", rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is down right now."+
		"<br><br><ul>"+
		"<br>The server, while acting as a gateway or proxy, received an invalid response from the upstream server it accessed in attempting to fulfill the request."+
		"<br>Please try again later and report the error to the website administrator"+
		"<br></ul>" {
		t.Errorf("expected body %q, got %q", "<br>The page you have requested is down right now."+
			"<br><br><ul>"+
			"<br>The server, while acting as a gateway or proxy, received an invalid response from the upstream server it accessed in attempting to fulfill the request."+
			"<br>Please try again later and report the error to the website administrator"+
			"<br></ul>", rw.Body.String())
	}
}
