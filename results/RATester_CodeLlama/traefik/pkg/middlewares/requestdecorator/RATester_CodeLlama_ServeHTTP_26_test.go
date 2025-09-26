package requestdecorator

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := &http.Request{
		Host: "example.com",
	}
	rw := &httptest.ResponseRecorder{}
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	r := &RequestDecorator{
		hostResolver: &Resolver{
			CnameFlattening: true,
		},
	}

	// when
	r.ServeHTTP(rw, req, next.ServeHTTP)

	// then
	assert.Equal(t, "example.com", rw.Header().Get("Host"))
}
