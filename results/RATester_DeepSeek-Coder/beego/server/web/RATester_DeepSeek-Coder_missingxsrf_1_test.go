package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMissingxsrf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("POST", "/some/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(missingxsrf)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}

	expected := "<br>The page you have requested is forbidden." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>'_xsrf' argument missing from POST" +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
