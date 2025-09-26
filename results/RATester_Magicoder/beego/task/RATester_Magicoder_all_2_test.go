package task

import (
	"fmt"
	"testing"
)

func Testall_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		bounds bounds
		want   uint64
	}

	tests := []test{
		{bounds{min: 0, max: 10, names: map[string]uint{"a": 1, "b": 2}}, 0x3FF},
		{bounds{min: 11, max: 20, names: map[string]uint{"a": 1, "b": 2}}, 0},
		{bounds{min: 21, max: 30, names: map[string]uint{"a": 1, "b": 2}}, 0x400},
	}

	for _, tt := range tests {
		got := all(tt.bounds)
		if got != tt.want {
			t.Errorf("all(%v) = %v, want %v", tt.bounds, got, tt.want)
		}
	}
}
