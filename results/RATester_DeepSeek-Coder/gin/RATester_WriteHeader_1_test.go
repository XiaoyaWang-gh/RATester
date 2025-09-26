package gin

import (
	"fmt"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	rw := &responseWriter{
		size:   0,
		status: 0,
	}

	testCases := []struct {
		name     string
		code     int
		expected int
	}{
		{"Test with 200", 200, 200},
		{"Test with 404", 404, 404},
		{"Test with 0", 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rw.WriteHeader(tc.code)
			if rw.status != tc.expected {
				t.Errorf("Expected status code %d, but got %d", tc.expected, rw.status)
			}
		})
	}
}
