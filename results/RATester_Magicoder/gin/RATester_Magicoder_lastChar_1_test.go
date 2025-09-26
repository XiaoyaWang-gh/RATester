package gin

import (
	"fmt"
	"testing"
)

func TestlastChar_1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want uint8
	}{
		{"Test Case 1", "Hello", 'o'},
		{"Test Case 2", "World", 'd'},
		{"Test Case 3", "Golang", 'g'},
		{"Test Case 4", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := lastChar(tt.str); got != tt.want {
				t.Errorf("lastChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
