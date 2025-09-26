package redirect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_31(t *testing.T) {
	t.Run("redirect", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		scheme := "https"
		hostname := "example.com"
		port := "8080"
		path := "/redirect"
		statusCode := http.StatusMovedPermanently

		r := redirect{
			scheme:     &scheme,
			hostname:   &hostname,
			port:       &port,
			path:       &path,
			statusCode: statusCode,
		}

		req, err := http.NewRequest("GET", "http://test.com/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != statusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, statusCode)
		}

		expectedLocation := fmt.Sprintf("%s://%s:%s%s", scheme, hostname, port, path)
		if location := rr.Header().Get("Location"); location != expectedLocation {
			t.Errorf("handler returned wrong location: got %v want %v", location, expectedLocation)
		}
	})
}
