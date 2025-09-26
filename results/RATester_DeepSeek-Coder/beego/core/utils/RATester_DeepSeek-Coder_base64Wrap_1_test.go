package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBase64Wrap_1(t *testing.T) {
	t.Run("Test base64Wrap function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    []byte
			expected string
		}{
			{
				name:     "Test case 1",
				input:    []byte("Hello, World!"),
				expected: "Hello, World!\r\n",
			},
			{
				name:     "Test case 2",
				input:    []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur."),
				expected: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur.\r\n",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				w := &bytes.Buffer{}
				base64Wrap(w, tc.input)
				result := w.String()
				if result != tc.expected {
					t.Errorf("Expected %q, got %q", tc.expected, result)
				}
			})
		}
	})
}
