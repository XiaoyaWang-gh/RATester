package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_2(t *testing.T) {
	length := Length{N: 5, Key: "test"}

	tests := []struct {
		name string
		obj  interface{}
		want bool
	}{
		{"string", "hello", true},
		{"string", "world", false},
		{"slice", []int{1, 2, 3, 4, 5}, true},
		{"slice", []int{1, 2, 3, 4}, false},
		{"not string or slice", 123, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := length.IsSatisfied(tt.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
