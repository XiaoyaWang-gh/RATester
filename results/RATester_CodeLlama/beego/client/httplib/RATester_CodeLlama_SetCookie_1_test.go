package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	cookie := &http.Cookie{}
	b.SetCookie(cookie)
	if b.req.Header.Get("Cookie") != cookie.String() {
		t.Errorf("SetCookie failed")
	}
}
