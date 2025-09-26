package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestforEachMediaRange_1(t *testing.T) {
	testCases := []struct {
		name     string
		header   []byte
		expected [][]byte
	}{
		{
			name:     "no quotes",
			header:   []byte("text/plain, text/html, text/css"),
			expected: [][]byte{[]byte("text/plain"), []byte("text/html"), []byte("text/css")},
		},
		{
			name:     "quotes",
			header:   []byte(`"text/plain", "text/html", "text/css"`),
			expected: [][]byte{[]byte("text/plain"), []byte("text/html"), []byte("text/css")},
		},
		{
			name:     "mixed",
			header:   []byte(`"text/plain", text/html, "text/css"`),
			expected: [][]byte{[]byte("text/plain"), []byte("text/html"), []byte("text/css")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var actual [][]byte
			forEachMediaRange(tc.header, func(b []byte) {
				actual = append(actual, b)
			})

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
