package compare

import (
	"fmt"
	"testing"
)

func TestLessStrings_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"Test1", "abc", "def", true},
		{"Test2", "def", "abc", false},
		{"Test3", "abc", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := LessStrings(tt.s, tt.t); got != tt.want {
				t.Errorf("LessStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
