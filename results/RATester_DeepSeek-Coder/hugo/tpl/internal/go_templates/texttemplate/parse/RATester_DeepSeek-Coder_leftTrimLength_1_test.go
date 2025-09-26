package parse

import (
	"fmt"
	"testing"
)

func TestLeftTrimLength_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Pos
	}{
		{"Test 1", "   hello", 3},
		{"Test 2", " world", 0},
		{"Test 3", "   ", 3},
		{"Test 4", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := leftTrimLength(tt.args); got != tt.want {
				t.Errorf("leftTrimLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
