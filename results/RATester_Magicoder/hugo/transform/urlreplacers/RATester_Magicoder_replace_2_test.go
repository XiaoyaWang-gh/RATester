package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func Testreplace_2(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected []byte
	}{
		{
			name:     "simple",
			content:  []byte("http://example.com"),
			expected: []byte("http://example.com"),
		},
		{
			name:     "with quotes",
			content:  []byte("http://example.com\"http://example.org"),
			expected: []byte("http://example.comhttp://example.org"),
		},
		{
			name:     "with quotes and spaces",
			content:  []byte("http://example.com \"http://example.org"),
			expected: []byte("http://example.com http://example.org"),
		},
		{
			name:     "with quotes and spaces and tabs",
			content:  []byte("http://example.com\t\"http://example.org"),
			expected: []byte("http://example.com http://example.org"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &absurllexer{
				content: test.content,
			}
			l.replace()
			if !bytes.Equal(l.content, test.expected) {
				t.Errorf("expected %s, got %s", test.expected, l.content)
			}
		})
	}
}
