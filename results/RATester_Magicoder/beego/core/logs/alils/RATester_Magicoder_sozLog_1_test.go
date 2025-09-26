package alils

import (
	"fmt"
	"testing"
)

func TestsozLog_1(t *testing.T) {
	tests := []struct {
		name string
		x    uint64
		want int
	}{
		{"Test 1", 1, 0},
		{"Test 2", 2, 1},
		{"Test 3", 3, 1},
		{"Test 4", 4, 2},
		{"Test 5", 5, 2},
		{"Test 6", 6, 2},
		{"Test 7", 7, 3},
		{"Test 8", 8, 3},
		{"Test 9", 9, 3},
		{"Test 10", 10, 3},
		{"Test 11", 11, 3},
		{"Test 12", 12, 4},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sozLog(tt.x); got != tt.want {
				t.Errorf("sozLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
