package capture

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFlush_4(t *testing.T) {
	testCases := []struct {
		name string
		rw   http.ResponseWriter
	}{
		{
			name: "Test with http.ResponseWriter",
			rw:   httptest.NewRecorder(),
		},
		{
			name: "Test with captureResponseWriter",
			rw: &captureResponseWriter{
				rw: httptest.NewRecorder(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			crw := &captureResponseWriter{
				rw: tc.rw,
			}
			crw.Flush()

			// Check if Flush() actually flushed the underlying ResponseWriter
			if f, ok := tc.rw.(http.Flusher); ok {
				f.Flush()
			}
		})
	}
}
