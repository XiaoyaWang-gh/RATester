package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(notFound)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	expected := "<br>The page you have requested has flown the coop." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>The page has moved" +
		"<br>The page no longer exists" +
		"<br>You were looking for your puppy and got lost" +
		"<br>You like 404 pages" +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
