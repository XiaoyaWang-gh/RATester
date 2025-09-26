package web

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestInvalidxsrf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	invalidxsrf(rw, r)
	if rw.Code != 417 {
		t.Errorf("invalidxsrf() = %v, want %v", rw.Code, 417)
	}
	if rw.Body.String() != "<br>The page you have requested is forbidden."+
		"<br>Perhaps you are here because:"+
		"<br><br><ul>"+
		"<br>expected XSRF not found"+
		"</ul>" {
		t.Errorf("invalidxsrf() = %v, want %v", rw.Body.String(), "<br>The page you have requested is forbidden."+
			"<br>Perhaps you are here because:"+
			"<br><br><ul>"+
			"<br>expected XSRF not found"+
			"</ul>")
	}
}
