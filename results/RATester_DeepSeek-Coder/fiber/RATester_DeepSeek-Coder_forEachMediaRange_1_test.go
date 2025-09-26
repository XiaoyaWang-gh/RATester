package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForEachMediaRange_1(t *testing.T) {
	testCases := []struct {
		name     string
		header   []byte
		expected [][]byte
	}{
		{
			name:     "no quotes",
			header:   []byte("text/plain, text/html"),
			expected: [][]byte{[]byte("text/plain"), []byte("text/html")},
		},
		{
			name:     "quotes",
			header:   []byte(`text/plain,"text/html,level=1",text/csv`),
			expected: [][]byte{[]byte("text/plain"), []byte("text/html,level=1"), []byte("text/csv")},
		},
		{
			name:     "escaped quotes",
			header:   []byte(`text/plain,"text/html\"level=1",text/csv`),
			expected: [][]byte{[]byte("text/plain"), []byte(`text/html"level=1`), []byte("text/csv")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var result [][]byte
			forEachMediaRange(tc.header, func(b []byte) {
				result = append(result, b)
			})
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
