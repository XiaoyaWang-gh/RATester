package render

import (
	"fmt"
	"io"
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
			callback: "callback1",
			data:     "data1",
			expected: "callback1({\"data\":\"data1\"})",
		},
		{
			name:     "Test case 2",
			callback: "callback2",
			data:     "data2",
			expected: "callback2({\"data\":\"data2\"})",
		},
		// Add more test cases as needed
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
				t.Errorf("Unexpected error: %v", err)
			}

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			if string(body) != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, string(body))
			}
		})
	}
}
