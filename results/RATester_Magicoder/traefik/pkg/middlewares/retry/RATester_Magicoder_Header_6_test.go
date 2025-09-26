package retry

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{
		headers: make(http.Header),
	}

	header := rw.Header()
	if header == nil {
		t.Error("Expected non-nil header")
	}

	rw.written = true
	header = rw.Header()
	if header == nil {
		t.Error("Expected non-nil header after written is set to true")
	}
}
