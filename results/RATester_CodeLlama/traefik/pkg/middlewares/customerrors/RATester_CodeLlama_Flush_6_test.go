package customerrors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFlush_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &codeModifier{
		code: http.StatusOK,
	}

	if flusher, ok := r.responseWriter.(http.Flusher); ok {
		flusher.Flush()
	}

	if r.headerSent {
		t.Errorf("headerSent = true, want false")
	}

	if r.code != http.StatusOK {
		t.Errorf("code = %d, want %d", r.code, http.StatusOK)
	}
}
