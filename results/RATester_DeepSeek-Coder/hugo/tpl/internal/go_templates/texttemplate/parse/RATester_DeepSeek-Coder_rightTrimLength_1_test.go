package parse

import (
	"fmt"
	"testing"
)

func TestRightTrimLength_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Pos
	}{
		{"Test1", "   hello   ", 5},
		{"Test2", "hello   ", 5},
		{"Test3", "   hello", 7},
		{"Test4", "hello", 5},
		{"Test5", "   ", 3},
		{"Test6", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := rightTrimLength(tt.args); got != tt.want {
				t.Errorf("rightTrimLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
