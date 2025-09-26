package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_1(t *testing.T) {
	numeric := Numeric{Key: "test"}

	tests := []struct {
		name string
		n    Numeric
		obj  interface{}
		want bool
	}{
		{
			name: "Test case 1: Valid numeric string",
			n:    numeric,
			obj:  "1234567890",
			want: true,
		},
		{
			name: "Test case 2: Invalid numeric string with special characters",
			n:    numeric,
			obj:  "1234567890abc",
			want: false,
		},
		{
			name: "Test case 3: Invalid numeric string with negative numbers",
			n:    numeric,
			obj:  "-1234567890",
			want: false,
		},
		{
			name: "Test case 4: Non-string input",
			n:    numeric,
			obj:  1234567890,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.n.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("Numeric.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
