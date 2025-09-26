package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_1(t *testing.T) {
	numeric := Numeric{Key: "test"}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{
			name: "Test case 1: Valid numeric string",
			obj:  "1234567890",
			want: true,
		},
		{
			name: "Test case 2: Invalid numeric string",
			obj:  "1234567890a",
			want: false,
		},
		{
			name: "Test case 3: Non-string input",
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

			if got := numeric.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
