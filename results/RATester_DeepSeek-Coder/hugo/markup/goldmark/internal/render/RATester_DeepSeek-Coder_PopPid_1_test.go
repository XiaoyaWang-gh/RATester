package render

import (
	"fmt"
	"testing"
)

func TestPopPid_1(t *testing.T) {
	ctx := &Context{
		pids: []uint64{1, 2, 3},
	}

	tests := []struct {
		name string
		want uint64
	}{
		{"Test 1", 3},
		{"Test 2", 2},
		{"Test 3", 1},
		{"Test 4", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ctx.PopPid(); got != tt.want {
				t.Errorf("PopPid() = %v, want %v", got, tt.want)
			}
		})
	}
}
