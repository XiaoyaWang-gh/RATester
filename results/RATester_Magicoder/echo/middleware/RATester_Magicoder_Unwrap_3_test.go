package middleware

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestUnwrap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &bodyDumpResponseWriter{
		ResponseWriter: httptest.NewRecorder(),
	}

	expected := httptest.NewRecorder()

	if result := rw.Unwrap(); result != expected {
		t.Errorf("Expected Unwrap to return %v, but got %v", expected, result)
	}
}
