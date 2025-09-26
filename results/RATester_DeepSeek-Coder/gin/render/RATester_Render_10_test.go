package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_10(t *testing.T) {
	testCases := []struct {
		name     string
		data     any
		expected string
	}{
		{
			name:     "Test with string data",
			data:     "test",
			expected: "test",
		},
		{
			name:     "Test with int data",
			data:     123,
			expected: "123",
		},
		{
			name:     "Test with special characters",
			data:     "😀",
			expected: "\\u1f600",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := httptest.NewRecorder()
			r := AsciiJSON{Data: tc.data}
			err := r.Render(w)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			result := w.Body.String()
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
