package capture

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	crw := &captureResponseWriter{rw: rw}

	status := http.StatusOK
	crw.WriteHeader(status)

	if crw.status != status {
		t.Errorf("Expected status %d, got %d", status, crw.status)
	}
}
