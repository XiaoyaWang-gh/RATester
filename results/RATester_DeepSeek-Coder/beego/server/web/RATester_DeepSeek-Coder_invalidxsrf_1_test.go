package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidxsrf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/invalidxsrf", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invalidxsrf)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusExpectationFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusExpectationFailed)
	}

	expected := "<br>The page you have requested is forbidden." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>expected XSRF not found" +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
