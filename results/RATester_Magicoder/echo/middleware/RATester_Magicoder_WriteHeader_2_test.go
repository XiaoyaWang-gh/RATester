package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &gzipResponseWriter{
		wroteHeader: false,
		code:        0,
	}

	w.WriteHeader(http.StatusOK)

	if !w.wroteHeader {
		t.Error("WriteHeader did not set wroteHeader to true")
	}

	if w.code != http.StatusOK {
		t.Errorf("WriteHeader did not set the correct status code, expected %d, got %d", http.StatusOK, w.code)
	}
}
