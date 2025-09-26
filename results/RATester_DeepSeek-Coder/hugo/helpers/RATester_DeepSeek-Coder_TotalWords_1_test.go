package helpers

import (
	"fmt"
	"testing"
)

func TestTotalWords_1(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{"", 0},
		{"Hello, world", 2},
		{"This is a test", 4},
		{"A sentence with multiple   spaces", 5},
		{"A sentence with    tabs\tand\tnewline\ncharacters", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := TotalWords(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
