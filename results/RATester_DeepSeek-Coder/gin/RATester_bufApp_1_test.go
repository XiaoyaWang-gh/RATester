package gin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBufApp_1(t *testing.T) {
	t.Run("Testing bufApp function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			buf      *[]byte
			s        string
			w        int
			c        byte
			expected []byte
		}{
			{
				name:     "Test case 1",
				buf:      new([]byte),
				s:        "hello",
				w:        1,
				c:        'a',
				expected: []byte("hallo"),
			},
			{
				name:     "Test case 2",
				buf:      new([]byte),
				s:        "world",
				w:        2,
				c:        'b',
				expected: []byte("wrbld"),
			},
			{
				name:     "Test case 3",
				buf:      new([]byte),
				s:        "go",
				w:        1,
				c:        'l',
				expected: []byte("gol"),
			},
			{
				name:     "Test case 4",
				buf:      new([]byte),
				s:        "lang",
				w:        4,
				c:        'u',
				expected: []byte("lang"),
			},
			{
				name:     "Test case 5",
				buf:      new([]byte),
				s:        "golang",
				w:        0,
				c:        'g',
				expected: []byte("golang"),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				bufApp(tc.buf, tc.s, tc.w, tc.c)
				if !bytes.Equal(*tc.buf, tc.expected) {
					t.Errorf("Expected %v, got %v", tc.expected, *tc.buf)
				}
			})
		}
	})
}
