package transform

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFrom_1(t *testing.T) {
	testCases := []struct {
		name     string
		from     *bytes.Buffer
		expected []byte
	}{
		{
			name:     "Test case 1",
			from:     bytes.NewBufferString("test"),
			expected: []byte("test"),
		},
		{
			name:     "Test case 2",
			from:     bytes.NewBufferString(""),
			expected: []byte(""),
		},
		{
			name:     "Test case 3",
			from:     bytes.NewBufferString("1234567890"),
			expected: []byte("1234567890"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ft := fromToBuffer{from: tc.from}
			result := ft.From().Bytes()
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
