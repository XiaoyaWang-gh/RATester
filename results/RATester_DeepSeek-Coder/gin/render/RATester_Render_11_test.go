package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_11(t *testing.T) {
	testCases := []struct {
		name     string
		callback string
		data     any
		expected string
	}{
		{
			name:     "Test case 1",
			callback: "callback",
			data:     "data",
			expected: "callback(data);",
		},
		{
			name:     "Test case 2",
			callback: "callback",
			data:     map[string]interface{}{"key": "value"},
			expected: "callback({\"key\":\"value\"});",
		},
		{
			name:     "Test case 3",
			callback: "callback",
			data:     struct{ Name string }{"John"},
			expected: "callback({\"Name\":\"John\"});",
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
			r := JsonpJSON{
				Callback: tc.callback,
				Data:     tc.data,
			}

			err := r.Render(w)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			result := w.Body.String()
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
