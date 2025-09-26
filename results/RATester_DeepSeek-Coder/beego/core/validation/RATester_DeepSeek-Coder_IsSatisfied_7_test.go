package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_7(t *testing.T) {
	a := Alpha{}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{"Test Case 1", "abc", true},
		{"Test Case 2", "123", false},
		{"Test Case 3", "abc123", false},
		{"Test Case 4", "ABC", true},
		{"Test Case 5", "AbC", true},
		{"Test Case 6", "abcABC", true},
		{"Test Case 7", "abcABC123", false},
		{"Test Case 8", "abcABC ", false},
		{"Test Case 9", " abcABC", false},
		{"Test Case 10", "abcABC@", false},
		{"Test Case 11", 123, false},
		{"Test Case 12", "@", false},
		{"Test Case 13", "", true},
		{"Test Case 14", " ", false},
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
