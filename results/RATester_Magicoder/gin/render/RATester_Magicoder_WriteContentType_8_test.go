package render

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestWriteContentType_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Redirect{
		Code: 301,
		Request: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   "example.com",
				Path:   "/path",
			},
		},
		Location: "http://example.com/newpath",
	}

	w := httptest.NewRecorder()
	r.WriteContentType(w)

	if w.Code != r.Code {
		t.Errorf("Expected status code %d, but got %d", r.Code, w.Code)
	}

	if w.Header().Get("Location") != r.Location {
		t.Errorf("Expected location header %s, but got %s", r.Location, w.Header().Get("Location"))
	}
}
