package urlreplacers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEmit_1(t *testing.T) {
	type testCase struct {
		name     string
		content  []byte
		expected []byte
		w        *bytes.Buffer
	}

	testCases := []testCase{
		{
			name:     "simple emit",
			content:  []byte("hello world"),
			expected: []byte("hello world"),
			w:        bytes.NewBuffer([]byte{}),
		},
		{
			name:     "emit with spaces",
			content:  []byte("hello  world"),
			expected: []byte("hello  world"),
			w:        bytes.NewBuffer([]byte{}),
		},
		{
			name:     "emit with special characters",
			content:  []byte("hello$%^&*()_+`~-=[]\\{}|;':\",.<>/?world"),
			expected: []byte("hello$%^&*()_+`~-=[]\\{}|;':\",.<>/?world"),
			w:        bytes.NewBuffer([]byte{}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &absurllexer{
				content: tc.content,
				w:       tc.w,
			}

			l.emit()

			if !bytes.Equal(tc.w.Bytes(), tc.expected) {
				t.Errorf("Expected %s, got %s", tc.expected, tc.w.Bytes())
			}
		})
	}
}
