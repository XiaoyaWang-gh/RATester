package render

import (
	"fmt"
	"testing"
)

func TestPushPid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		pids: []uint64{},
	}

	pid := uint64(123)
	ctx.PushPid(pid)

	if len(ctx.pids) != 1 {
		t.Errorf("Expected length of pids to be 1, but got %d", len(ctx.pids))
	}

	if ctx.pids[0] != pid {
		t.Errorf("Expected pid to be %d, but got %d", pid, ctx.pids[0])
	}
}
