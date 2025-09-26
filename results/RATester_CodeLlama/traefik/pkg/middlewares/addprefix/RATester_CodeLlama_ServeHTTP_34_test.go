package addprefix

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_34(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	a := &addPrefix{
		next:   http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {}),
		prefix: "/prefix",
	}
	a.ServeHTTP(rw, req)
}
