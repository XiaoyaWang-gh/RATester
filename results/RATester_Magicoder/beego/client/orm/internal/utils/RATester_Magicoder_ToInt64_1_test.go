package utils

import (
	"fmt"
	"testing"
)

func TestToInt64_1(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  int64
	}{
		{"int", 1, 1},
		{"int8", int8(2), 2},
		{"int16", int16(3), 3},
		{"int32", int32(4), 4},
		{"int64", int64(5), 5},
		{"uint", uint(6), 6},
		{"uint8", uint8(7), 7},
		{"uint16", uint16(8), 8},
		{"uint32", uint32(9), 9},
		{"uint64", uint64(10), 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToInt64(tt.value); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
