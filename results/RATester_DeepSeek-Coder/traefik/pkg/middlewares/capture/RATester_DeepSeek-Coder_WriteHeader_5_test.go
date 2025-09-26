package capture

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader_5(t *testing.T) {
	testCases := []struct {
		name   string
		status int
	}{
		{"200 OK", http.StatusOK},
		{"404 Not Found", http.StatusNotFound},
		{"500 Internal Server Error", http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := httptest.NewRecorder()
			crw := &captureResponseWriter{rw: rw}

			crw.WriteHeader(tc.status)

			if crw.status != tc.status {
				t.Errorf("Expected status %d, got %d", tc.status, crw.status)
			}

			if rw.Code != tc.status {
				t.Errorf("Expected status %d, got %d", tc.status, rw.Code)
			}
		})
	}
}
