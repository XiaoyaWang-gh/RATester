package collections

import (
	"fmt"
	"testing"
)

func TestCount_3(t *testing.T) {
	ss := SortedStringSlice{"apple", "banana", "apple", "cherry", "apple"}

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"apple", "apple", 3},
		{"banana", "banana", 1},
		{"cherry", "cherry", 1},
		{"not found", "not found", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ss.Count(tt.input); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
