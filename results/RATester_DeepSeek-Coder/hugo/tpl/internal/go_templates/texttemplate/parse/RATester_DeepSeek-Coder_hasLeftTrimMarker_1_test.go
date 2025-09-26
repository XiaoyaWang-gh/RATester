package parse

import (
	"fmt"
	"testing"
)

func TestHasLeftTrimMarker_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Test case 1", "  hello", true},
		{"Test case 2", " hello", false},
		{"Test case 3", "hello", false},
		{"Test case 4", "  hello world", true},
		{"Test case 5", "   hello world", true},
		{"Test case 6", "    hello world", true},
		{"Test case 7", "     hello world", true},
		{"Test case 8", "      hello world", true},
		{"Test case 9", "       hello world", true},
		{"Test case 10", "        hello world", true},
		{"Test case 11", "         hello world", true},
		{"Test case 12", "          hello world", true},
		{"Test case 13", "           hello world", true},
		{"Test case 14", "            hello world", true},
		{"Test case 15", "             hello world", true},
		{"Test case 16", "              hello world", true},
		{"Test case 17", "               hello world", true},
		{"Test case 18", "                hello world", true},
		{"Test case 19", "                 hello world", true},
		{"Test case 20", "                  hello world", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := hasLeftTrimMarker(tt.args); got != tt.want {
				t.Errorf("hasLeftTrimMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
