package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForbidden_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/forbidden", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(forbidden)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusForbidden)
	}

	expected := "<br>The page you have requested is forbidden." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>Your address may be blocked" +
		"<br>The site may be disabled" +
		"<br>You need to log in" +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
