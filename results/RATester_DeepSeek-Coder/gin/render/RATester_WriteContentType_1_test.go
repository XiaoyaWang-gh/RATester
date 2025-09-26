package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestWriteContentType_1(t *testing.T) {
	testCases := []struct {
		name     string
		data     Data
		expected string
	}{
		{
			name: "Test case 1",
			data: Data{
				ContentType: "application/json",
				Data:        []byte("test"),
			},
			expected: "application/json",
		},
		{
			name: "Test case 2",
			data: Data{
				ContentType: "application/xml",
				Data:        []byte("test"),
			},
			expected: "application/xml",
		},
		{
			name: "Test case 3",
			data: Data{
				ContentType: "text/plain",
				Data:        []byte("test"),
			},
			expected: "text/plain",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			recorder := httptest.NewRecorder()
			tc.data.WriteContentType(recorder)
			result := recorder.Header().Get("Content-Type")
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
