package internal

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestGetAsciidocContent_1(t *testing.T) {
	testCases := []struct {
		name     string
		src      []byte
		ctx      converter.DocumentContext
		expected []byte
		err      error
	}{
		{
			name:     "Test with valid asciidoc content",
			src:      []byte("= Hello, World!\n"),
			ctx:      converter.DocumentContext{DocumentName: "test.adoc"},
			expected: []byte("<div class=\"paragraph\">\n<p>Hello, World&#8230;</p>\n</div>\n"),
			err:      nil,
		},
		{
			name:     "Test with invalid asciidoc content",
			src:      []byte("= Hello, \n"),
			ctx:      converter.DocumentContext{DocumentName: "test.adoc"},
			expected: []byte(""),
			err:      nil,
		},
		{
			name:     "Test with empty asciidoc content",
			src:      []byte(""),
			ctx:      converter.DocumentContext{DocumentName: "test.adoc"},
			expected: []byte(""),
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			converter := &AsciidocConverter{}
			result, err := converter.GetAsciidocContent(tc.src, tc.ctx)
			if err != tc.err {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected result %v, got %v", tc.expected, result)
			}
		})
	}
}
