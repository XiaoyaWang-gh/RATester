package hash

import (
	"fmt"
	"testing"
)

func TestFNV32a_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		testCases := []struct {
			name     string
			input    any
			expected int
		}{
			{
				name:     "string input",
				input:    "test",
				expected: 2811947799,
			},
			{
				name:     "int input",
				input:    123,
				expected: 3318065099,
			},
			{
				name:     "float input",
				input:    123.456,
				expected: 3318065099,
			},
			{
				name:     "bool input",
				input:    true,
				expected: 2166136279,
			},
		}

		for _, tc := range testCases {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				t.Parallel()

				result, err := ns.FNV32a(tc.input)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if result != tc.expected {
					t.Errorf("expected %v, got %v", tc.expected, result)
				}
			})
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		_, err := ns.FNV32a(make(chan int))
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
