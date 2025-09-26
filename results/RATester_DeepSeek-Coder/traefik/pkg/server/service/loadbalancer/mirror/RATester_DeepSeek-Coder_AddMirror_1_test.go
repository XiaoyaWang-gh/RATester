package mirror

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestAddMirror_1(t *testing.T) {
	type testCase struct {
		name     string
		handler  http.Handler
		percent  int
		expected error
	}

	testCases := []testCase{
		{
			name:     "Test with valid handler and percent",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			percent:  50,
			expected: nil,
		},
		{
			name:     "Test with valid handler and percent 100",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			percent:  100,
			expected: nil,
		},
		{
			name:     "Test with valid handler and percent 0",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			percent:  0,
			expected: nil,
		},
		{
			name:     "Test with invalid percent -1",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			percent:  -1,
			expected: errors.New("percent must be between 0 and 100"),
		},
		{
			name:     "Test with invalid percent 101",
			handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			percent:  101,
			expected: errors.New("percent must be between 0 and 100"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Mirroring{}
			err := m.AddMirror(tc.handler, tc.percent)
			if err != tc.expected {
				t.Errorf("Expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
