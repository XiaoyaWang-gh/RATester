package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInternalServerError_1(t *testing.T) {
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
	internalServerError(rw, r)
	if rw.Code != http.StatusInternalServerError {
		t.Errorf("expected %d, got %d", http.StatusInternalServerError, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is down right now."+
		"<br><br><ul>"+
		"<br>Please try again later and report the error to the website administrator"+
		"<br></ul>" {
		t.Errorf("expected %s, got %s", "<br>The page you have requested is down right now."+
			"<br><br><ul>"+
			"<br>Please try again later and report the error to the website administrator"+
			"<br></ul>", rw.Body.String())
	}
}
