package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGET_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	path := "/test"
	handler := func(c Context) error {
		return c.String(http.StatusOK, "test")
	}

	e.GET(path, handler)

	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "test" {
		t.Errorf("Expected body 'test', got '%s'", rec.Body.String())
	}
}
