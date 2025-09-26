package replacepath

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServeHTTP_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add(ReplacedPathHeader, "/")
	req.URL.RawPath = "/"
	req.URL.Path, _ = url.PathUnescape(req.URL.RawPath)
	req.RequestURI = req.URL.RequestURI()

	r := &replacePath{
		next: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/", req.URL.Path)
		}),
		path: "/",
		name: "test",
	}

	r.ServeHTTP(rw, req)
}
