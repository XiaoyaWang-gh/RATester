package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_5(t *testing.T) {
	an := AlphaNumeric{}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{"Test Case 1", "abc123", true},
		{"Test Case 2", "abc@123", false},
		{"Test Case 3", "abc_123", false},
		{"Test Case 4", "abc-123", false},
		{"Test Case 5", "abc123_", false},
		{"Test Case 6", "abc123-", false},
		{"Test Case 7", "abc123 ", false},
		{"Test Case 8", " abc123", false},
		{"Test Case 9", 123456, false},
		{"Test Case 10", "123456", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := an.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
