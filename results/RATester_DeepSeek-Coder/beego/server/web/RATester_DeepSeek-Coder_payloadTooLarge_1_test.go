package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPayloadTooLarge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("POST", "/payloadTooLarge", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(payloadTooLarge)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusRequestEntityTooLarge {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusRequestEntityTooLarge)
	}

	expected := `<br>The page you have requested is unavailable.
		 <br>Perhaps you are here because:<br><br>
		 <ul>
			<br>The request entity is larger than limits defined by server.
			<br>Please change the request entity and try again.
		 </ul>
		`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
