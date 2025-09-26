package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	s := &traceablePlugin{
		name: "test",
		h:    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	s.ServeHTTP(rw, req)
}
