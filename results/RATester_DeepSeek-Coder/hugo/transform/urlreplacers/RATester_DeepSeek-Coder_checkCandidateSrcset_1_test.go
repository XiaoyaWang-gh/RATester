package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCheckCandidateSrcset_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected []byte
	}{
		{
			name:     "simple case",
			content:  []byte(`srcset="/path/to/image.jpg"`),
			expected: []byte(`/path/to/image.jpg`),
		},
		{
			name:     "multiple images",
			content:  []byte(`srcset="/path/to/image1.jpg 1x, /path/to/image2.jpg 2x"`),
			expected: []byte(`/path/to/image1.jpg 1x, /path/to/image2.jpg 2x`),
		},
		{
			name:     "absolute url",
			content:  []byte(`srcset="https://example.com/path/to/image.jpg"`),
			expected: []byte(`https://example.com/path/to/image.jpg`),
		},
		{
			name:     "relative url",
			content:  []byte(`srcset="./path/to/image.jpg"`),
			expected: []byte(`./path/to/image.jpg`),
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
			}

			checkCandidateSrcset(l)

			got := l.w.(*bytes.Buffer).Bytes()
			if !bytes.Equal(got, tt.expected) {
				t.Errorf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
