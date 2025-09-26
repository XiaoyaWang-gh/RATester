package strings

import (
	"errors"
	"fmt"
	"html/template"
	"testing"
)

func TestChomp_2(t *testing.T) {
	ns := &Namespace{}

	t.Run("Chomp", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			name     string
			s        any
			expected any
			err      error
		}{
			{
				name:     "Chomp string",
				s:        "test\n",
				expected: "test",
				err:      nil,
			},
			{
				name:     "Chomp HTML",
				s:        template.HTML("test\n"),
				expected: template.HTML("test"),
				err:      nil,
			},
			{
				name:     "Chomp non-string",
				s:        123,
				expected: "",
				err:      errors.New("cast: unable to cast 123 of type int to string"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				actual, err := ns.Chomp(tt.s)
				if err != nil && tt.err == nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
					t.Errorf("Expected error %v, got %v", tt.err, err)
				}
				if actual != tt.expected {
					t.Errorf("Expected %v, got %v", tt.expected, actual)
				}
			})
		}
	})
}
