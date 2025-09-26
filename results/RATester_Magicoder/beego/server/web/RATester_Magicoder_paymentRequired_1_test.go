package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestpaymentRequired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req, err := http.NewRequest("GET", "/paymentRequired", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(paymentRequired)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPaymentRequired {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPaymentRequired)
	}

	expected := "<br>The page you have requested Payment Required." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>The credentials you supplied are incorrect" +
		"<br>There are errors in the website address" +
		"</ul>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
