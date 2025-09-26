package glob

import (
	"fmt"
	"testing"
)

func TestHasGlobChar_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test1", "*", true},
		{"Test2", "abc", false},
		{"Test3", "a*c", true},
		{"Test4", "a**c", true},
		{"Test5", "a?c", true},
		{"Test6", "a**?c", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasGlobChar(tt.arg); got != tt.want {
				t.Errorf("HasGlobChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
