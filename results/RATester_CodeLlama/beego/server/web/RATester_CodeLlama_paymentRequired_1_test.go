package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentRequired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	paymentRequired(rw, r)
	if rw.Code != http.StatusPaymentRequired {
		t.Errorf("Expected %v, got %v", http.StatusPaymentRequired, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested Payment Required."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>The credentials you supplied are incorrect"+
		"<br>There are errors in the website address"+
		"</ul>" {
		t.Errorf("Expected %v, got %v", "<br>The page you have requested Payment Required."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>The credentials you supplied are incorrect"+
			"<br>There are errors in the website address"+
			"</ul>", rw.Body.String())
	}
}
