package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceUnavailable_1(t *testing.T) {
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
	serviceUnavailable(rw, r)
	if rw.Code != http.StatusServiceUnavailable {
		t.Errorf("expected %d, got %d", http.StatusServiceUnavailable, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is unavailable."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br><br>The page is overloaded"+
		"<br>Please try again later."+
		"</ul>" {
		t.Errorf("expected %q, got %q", "<br>The page you have requested is unavailable."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br><br>The page is overloaded"+
			"<br>Please try again later."+
			"</ul>", rw.Body.String())
	}
}
