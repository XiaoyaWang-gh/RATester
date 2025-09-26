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
		pids: make([]uint64, 0),
	}

	pid := uint64(123456789)
	ctx.PushPid(pid)

	if len(ctx.pids) != 1 {
		t.Errorf("Expected length of pids slice to be 1, got %d", len(ctx.pids))
	}

	if ctx.pids[0] != pid {
		t.Errorf("Expected first element of pids slice to be %d, got %d", pid, ctx.pids[0])
	}
}
