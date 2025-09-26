package alils

import (
	"fmt"
	"testing"
)

func TestsovLog_1(t *testing.T) {
	tests := []struct {
		name string
		x    uint64
		want int
	}{
		{"Test 1", 0, 0},
		{"Test 2", 1, 1},
		{"Test 3", 128, 2},
		{"Test 4", 1024, 3},
		{"Test 5", 1000000, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sovLog(tt.x); got != tt.want {
				t.Errorf("sovLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
