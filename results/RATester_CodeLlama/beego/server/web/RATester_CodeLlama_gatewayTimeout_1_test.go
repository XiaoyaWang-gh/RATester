package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGatewayTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// and: A ResponseRecorder
	rr := httptest.NewRecorder()

	// when: We call the method under test
	gatewayTimeout(rr, req)

	// then: We expect a 504 status code
	if status := rr.Code; status != http.StatusGatewayTimeout {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusGatewayTimeout)
	}

	// and: We expect a response body
	expected := "<br>The page you have requested is unavailable" +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br><br>The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server specified by the URI." +
		"<br>Please try again later." +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
