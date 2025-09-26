package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed_1(t *testing.T) {
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
	handler := http.HandlerFunc(methodNotAllowed)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	expected := "<br>The method you have requested Not Allowed." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>The method specified in the Request-Line is not allowed for the resource identified by the Request-URI" +
		"<br>The response MUST include an Allow header containing a list of valid methods for the requested resource." +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
