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

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	notFound(rw, r)
	if rw.Code != http.StatusNotFound {
		t.Errorf("Expected %d, got %d", http.StatusNotFound, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested has flown the coop."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>The page has moved"+
		"<br>The page no longer exists"+
		"<br>You were looking for your puppy and got lost"+
		"<br>You like 404 pages"+
		"</ul>" {
		t.Errorf("Expected %s, got %s", "<br>The page you have requested has flown the coop."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>The page has moved"+
			"<br>The page no longer exists"+
			"<br>You were looking for your puppy and got lost"+
			"<br>You like 404 pages", rw.Body.String())
	}
}
