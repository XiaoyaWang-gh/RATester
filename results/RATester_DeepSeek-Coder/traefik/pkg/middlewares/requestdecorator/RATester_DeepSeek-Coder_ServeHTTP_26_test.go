package requestdecorator

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestServeHTTP_26(t *testing.T) {
	t.Run("Test ServeHTTP", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &RequestDecorator{
				hostResolver: &Resolver{
					CnameFlattening: true,
					ResolvConfig:    "test",
					ResolvDepth:     1,
					cache:           &cache.Cache{},
				},
			}
			rw.ServeHTTP(w, r, func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, client")
			})
		})

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := "Hello, client"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
