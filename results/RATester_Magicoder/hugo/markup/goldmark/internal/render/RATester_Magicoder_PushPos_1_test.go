package render

import (
	"fmt"
	"testing"
)

func TestPushPos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		positions: []int{},
	}

	ctx.PushPos(1)
	ctx.PushPos(2)
	ctx.PushPos(3)

	if len(ctx.positions) != 3 {
		t.Errorf("Expected length of positions to be 3, but got %d", len(ctx.positions))
	}

	if ctx.positions[0] != 1 {
		t.Errorf("Expected first position to be 1, but got %d", ctx.positions[0])
	}

	if ctx.positions[1] != 2 {
		t.Errorf("Expected second position to be 2, but got %d", ctx.positions[1])
	}

	if ctx.positions[2] != 3 {
		t.Errorf("Expected third position to be 3, but got %d", ctx.positions[2])
	}
}
