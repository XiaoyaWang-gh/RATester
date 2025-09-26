package contenttype

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := &contentType{
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if ct, ok := rw.Header()["Content-Type"]; ok && ct == nil {
				t.Error("Content-Type should be deleted")
			}
		}),
	}
	c.ServeHTTP(rw, req)
}
