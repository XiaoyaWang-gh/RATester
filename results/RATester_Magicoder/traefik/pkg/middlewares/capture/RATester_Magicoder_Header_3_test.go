package capture

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{
		rw:     httptest.NewRecorder(),
		status: http.StatusOK,
	}

	header := crw.Header()

	if header == nil {
		t.Error("Expected non-nil header, got nil")
	}
}
