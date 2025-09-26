package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_3(t *testing.T) {
	testCases := []struct {
		name     string
		r        SecureJSON
		expected string
	}{
		{
			name: "Test case 1",
			r: SecureJSON{
				Prefix: "prefix",
				Data:   []string{"data1", "data2"},
			},
			expected: "prefix[data1,data2]",
		},
		{
			name: "Test case 2",
			r: SecureJSON{
				Prefix: "prefix",
				Data:   "data",
			},
			expected: "prefixdata",
		},
		{
			name: "Test case 3",
			r: SecureJSON{
				Prefix: "prefix",
				Data:   nil,
			},
			expected: "prefixnull",
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
			err := tc.r.Render(w)
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
