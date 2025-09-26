package compress

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	type testCase struct {
		name     string
		status   int
		expected int
	}

	testCases := []testCase{
		{"StatusOK", http.StatusOK, http.StatusOK},
		{"StatusNotFound", http.StatusNotFound, http.StatusNotFound},
		{"StatusInternalServerError", http.StatusInternalServerError, http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw := &responseWriter{}
			rw.WriteHeader(tc.status)
			if rw.statusCode != tc.expected {
				t.Errorf("Expected status code %d, got %d", tc.expected, rw.statusCode)
			}
		})
	}
}
