package capture

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRead_3(t *testing.T) {
	testCases := []struct {
		name     string
		source   io.ReadCloser
		expected int64
	}{
		{
			name:     "Test case 1",
			source:   ioutil.NopCloser(strings.NewReader("test")),
			expected: 4,
		},
		{
			name:     "Test case 2",
			source:   ioutil.NopCloser(strings.NewReader("")),
			expected: 0,
		},
		{
			name:     "Test case 3",
			source:   ioutil.NopCloser(strings.NewReader("Hello, World")),
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			counter := &readCounter{source: tc.source}
			buf := make([]byte, 1024)
			for {
				n, err := counter.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}
					t.Errorf("Unexpected error: %v", err)
				}
				if n == 0 {
					t.Errorf("Read 0 bytes")
				}
			}
			if counter.size != tc.expected {
				t.Errorf("Expected %d bytes, got %d", tc.expected, counter.size)
			}
		})
	}
}
