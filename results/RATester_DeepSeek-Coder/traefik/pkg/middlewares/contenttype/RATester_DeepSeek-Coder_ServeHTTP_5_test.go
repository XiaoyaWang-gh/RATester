package contenttype

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_5(t *testing.T) {
	t.Run("Test Content-Type auto-detection", func(t *testing.T) {

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
			w.Write([]byte("OK"))
		})

		ct := &contentType{
			next: handler,
			name: "test",
		}

		ct.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, rr.Code)
		}

		if rr.Header().Get("Content-Type") != "" {
			t.Errorf("Expected Content-Type to be empty, got %v", rr.Header().Get("Content-Type"))
		}
	})
}
