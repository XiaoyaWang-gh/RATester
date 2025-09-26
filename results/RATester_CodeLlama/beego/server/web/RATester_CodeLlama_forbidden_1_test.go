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

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	forbidden(rw, r)
	if rw.Code != http.StatusForbidden {
		t.Errorf("Expected %v, got %v", http.StatusForbidden, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is forbidden."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>Your address may be blocked"+
		"<br>The site may be disabled"+
		"<br>You need to log in"+
		"</ul>" {
		t.Errorf("Expected %v, got %v", "<br>The page you have requested is forbidden."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>Your address may be blocked"+
			"<br>The site may be disabled"+
			"<br>You need to log in"+
			"</ul>", rw.Body.String())
	}
}
