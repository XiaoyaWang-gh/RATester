package render

import (
	"fmt"
	"testing"
)

func TestPeekPid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{}
	ctx.pids = []uint64{1, 2, 3}
	if got := ctx.PeekPid(); got != 3 {
		t.Errorf("PeekPid() = %v, want %v", got, 3)
	}
}
