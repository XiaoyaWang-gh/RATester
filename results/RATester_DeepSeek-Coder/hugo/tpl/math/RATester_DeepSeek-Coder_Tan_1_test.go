package math

import (
	"fmt"
	"math"
	"testing"
)

func TestTan_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		tests := []struct {
			name     string
			input    any
			expected float64
		}{
			{"positive", 1, math.Tan(1)},
			{"negative", -1, math.Tan(-1)},
			{"zero", 0, 0},
		}

		for _, test := range tests {
			test := test
			t.Run(test.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				t.Parallel()

				result, err := ns.Tan(test.input)
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if result != test.expected {
					t.Errorf("expected %v, got %v", test.expected, result)
				}
			})
		}
	})

	t.Run("invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		_, err := ns.Tan("not a number")
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})
}
