package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTRACE_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()

	testPath := "/test"
	testHandler := func(c Context) error {
		return c.String(http.StatusOK, "test")
	}

	e.TRACE(testPath, testHandler)

	req := httptest.NewRequest(http.MethodTrace, testPath, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "test" {
		t.Errorf("Expected body 'test', got '%s'", rec.Body.String())
	}
}
