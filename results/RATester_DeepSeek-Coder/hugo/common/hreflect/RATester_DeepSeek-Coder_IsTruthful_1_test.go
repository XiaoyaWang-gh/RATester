package hreflect

import (
	"fmt"
	"testing"
)

func TestIsTruthful_1(t *testing.T) {
	t.Run("TestIsTruthfulValue", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name string
			in   any
			want bool
		}{
			{
				name: "Test case 1",
				in:   "test",
				want: true,
			},
			{
				name: "Test case 2",
				in:   0,
				want: false,
			},
			{
				name: "Test case 3",
				in:   "",
				want: false,
			},
			{
				name: "Test case 4",
				in:   nil,
				want: false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				got := IsTruthful(tc.in)
				if got != tc.want {
					t.Errorf("IsTruthful(%v) = %v, want %v", tc.in, got, tc.want)
				}
			})
		}
	})
}
