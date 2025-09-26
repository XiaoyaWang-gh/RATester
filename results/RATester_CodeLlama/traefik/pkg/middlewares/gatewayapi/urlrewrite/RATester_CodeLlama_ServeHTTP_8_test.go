package urlrewrite

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	u := urlRewrite{
		name: "test",
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			t.Log("next handler")
		}),
	}
	u.ServeHTTP(rw, req)
}
