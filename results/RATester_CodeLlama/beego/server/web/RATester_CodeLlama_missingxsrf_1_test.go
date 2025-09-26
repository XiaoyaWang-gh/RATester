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

	rw := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	missingxsrf(rw, r)
	if rw.Code != 422 {
		t.Errorf("expected status code 422, got %d", rw.Code)
	}
	if rw.Body.String() != "<br>The page you have requested is forbidden."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>'_xsrf' argument missing from POST"+
		"</ul>" {
		t.Errorf("expected body %q, got %q", "<br>The page you have requested is forbidden."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>'_xsrf' argument missing from POST"+
			"</ul>", rw.Body.String())
	}
}
