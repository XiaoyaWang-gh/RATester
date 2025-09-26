package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatic_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	e.Static("/static", "./public")

	req := httptest.NewRequest(http.MethodGet, "/static/test.txt", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if rec.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected Content-Type header %q, got %q", expectedContentType, rec.Header().Get("Content-Type"))
	}

	expectedBody := "This is a test file."
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, rec.Body.String())
	}
}
