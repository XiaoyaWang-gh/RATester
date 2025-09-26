package gin

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestUnwrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{
		ResponseWriter: httptest.NewRecorder(),
		size:           0,
		status:         0,
	}

	expected := httptest.NewRecorder()

	result := rw.Unwrap()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
