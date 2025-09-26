package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCheckCandidateBase_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected []byte
	}{
		{
			name:     "test1",
			content:  []byte("http://example.com/path"),
			expected: []byte("http://example.com/path"),
		},
		{
			name:     "test2",
			content:  []byte("http://example.com/path/"),
			expected: []byte("http://example.com/path/"),
		},
		{
			name:     "test3",
			content:  []byte("http://example.com/path/subpath"),
			expected: []byte("http://example.com/path/subpath"),
		},
		{
			name:     "test4",
			content:  []byte("http://example.com/path/subpath/"),
			expected: []byte("http://example.com/path/subpath/"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &absurllexer{
				content: tt.content,
				w:       &bytes.Buffer{},
				path:    []byte("."),
			}

			checkCandidateBase(l)

			got := l.w.(*bytes.Buffer).Bytes()
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
