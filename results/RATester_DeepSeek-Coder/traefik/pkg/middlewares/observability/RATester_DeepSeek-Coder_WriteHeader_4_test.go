package observability

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteHeader_4(t *testing.T) {
	testCases := []struct {
		name   string
		status int
		want   int
	}{
		{"200", 200, 200},
		{"404", 404, 404},
		{"500", 500, 500},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			recorder := &statusCodeRecorder{
				ResponseWriter: httptest.NewRecorder(),
			}

			recorder.WriteHeader(tc.status)

			if got := recorder.Status(); got != tc.want {
				t.Errorf("Expected status code %d, got %d", tc.want, got)
			}
		})
	}
}
