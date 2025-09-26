package render

import (
	"fmt"
	"testing"
)

func TestPopPid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		pids: []uint64{1, 2, 3},
	}
	pid := ctx.PopPid()
	if pid != 3 {
		t.Errorf("Expected pid 3, got %d", pid)
	}
	if len(ctx.pids) != 2 {
		t.Errorf("Expected pids length 2, got %d", len(ctx.pids))
	}
}
