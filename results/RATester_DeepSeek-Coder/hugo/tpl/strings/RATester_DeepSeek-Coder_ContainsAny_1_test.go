package strings

import (
	"errors"
	"fmt"
	"testing"
)

func TestContainsAny_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("ContainsAny", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			name     string
			s        any
			chars    any
			expected bool
			err      error
		}{
			{
				name:     "ContainsAny",
				s:        "Hello, World",
				chars:    "e,",
				expected: true,
				err:      nil,
			},
			{
				name:     "ContainsAny",
				s:        "Hello, World",
				chars:    "z",
				expected: false,
				err:      nil,
			},
			{
				name:     "ContainsAny",
				s:        12345,
				chars:    "23",
				expected: true,
				err:      nil,
			},
			{
				name:     "ContainsAny",
				s:        "Hello, World",
				chars:    12345,
				expected: false,
				err:      errors.New("cast: unable to cast \"12345\" of type int to string"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				actual, err := ns.ContainsAny(tt.s, tt.chars)
				if err != nil && err.Error() != tt.err.Error() {
					t.Errorf("expected error %v, got %v", tt.err, err)
				}
				if actual != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, actual)
				}
			})
		}
	})
}
