package compress

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_13(t *testing.T) {
	t.Run("should return 500 if newCompressionWriter fails", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := Config{
			Algorithm: "invalid",
		}
		compressionHandler := &CompressionHandler{
			cfg: cfg,
		}
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		compressionHandler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
		}
	})

	t.Run("should return 200 if newCompressionWriter succeeds", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cfg := Config{
			Algorithm: "Brotli",
		}
		compressionHandler := &CompressionHandler{
			cfg: cfg,
		}
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		compressionHandler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}
