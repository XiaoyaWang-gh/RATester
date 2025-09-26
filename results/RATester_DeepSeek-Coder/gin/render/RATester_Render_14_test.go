package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_14(t *testing.T) {
	testCases := []struct {
		name     string
		data     any
		expected string
	}{
		{
			name:     "Test case 1",
			data:     "test",
			expected: "\"test\"\n",
		},
		{
			name:     "Test case 2",
			data:     123,
			expected: "123\n",
		},
		{
			name:     "Test case 3",
			data:     true,
			expected: "true\n",
		},
		{
			name:     "Test case 4",
			data:     nil,
			expected: "null\n",
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
			r := PureJSON{Data: tc.data}
			err := r.Render(w)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if w.Body.String() != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, w.Body.String())
			}
		})
	}
}
