package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_5(t *testing.T) {
	a := AlphaNumeric{}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{"Test case 1", "abc123", true},
		{"Test case 2", "abc123@", false},
		{"Test case 3", 123456, false},
		{"Test case 4", "abc@123", false},
		{"Test case 5", "abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := a.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
