package retry

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestnewResponseWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	shouldRetry := true
	r := newResponseWriter(rw, shouldRetry)

	if r.responseWriter != rw {
		t.Errorf("Expected responseWriter to be %v, but got %v", rw, r.responseWriter)
	}

	if r.headers == nil {
		t.Errorf("Expected headers to be initialized, but got nil")
	}

	if r.shouldRetry != shouldRetry {
		t.Errorf("Expected shouldRetry to be %v, but got %v", shouldRetry, r.shouldRetry)
	}

	if r.written {
		t.Errorf("Expected written to be false, but got true")
	}
}
