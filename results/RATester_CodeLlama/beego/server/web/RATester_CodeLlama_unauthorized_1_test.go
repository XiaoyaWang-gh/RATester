package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnauthorized_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	unauthorized(rw, r)
	if rw.Code != http.StatusUnauthorized {
		t.Errorf("Status code is not %d, but %d", http.StatusUnauthorized, rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested can't be authorized."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>The credentials you supplied are incorrect"+
		"<br>There are errors in the website address"+
		"</ul>" {
		t.Errorf("Body is not %s, but %s", "<br>The page you have requested can't be authorized."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>The credentials you supplied are incorrect"+
			"<br>There are errors in the website address"+
			"</ul>", rw.Body.String())
	}
}
