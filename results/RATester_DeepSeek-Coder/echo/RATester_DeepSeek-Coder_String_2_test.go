package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testString := "test string"
	err := c.String(http.StatusOK, testString)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != testString {
		t.Errorf("Expected body %q, got %q", testString, rec.Body.String())
	}
}
