package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestconsumeQuote_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		quotes   [][]byte
		expected []byte
	}{
		{
			name:     "test1",
			content:  []byte("\"test\""),
			quotes:   [][]byte{[]byte("\"")},
			expected: []byte("\"test\""),
		},
		{
			name:     "test2",
			content:  []byte("'test'"),
			quotes:   [][]byte{[]byte("'")},
			expected: []byte("'test'"),
		},
		{
			name:     "test3",
			content:  []byte("`test`"),
			quotes:   [][]byte{[]byte("`")},
			expected: []byte("`test`"),
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
				quotes:  test.quotes,
			}
			result := l.consumeQuote()
			if !bytes.Equal(result, test.expected) {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
