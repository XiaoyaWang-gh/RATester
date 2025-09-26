package render

import (
	"fmt"
	"testing"
)

func TestPopPos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		positions: []int{1, 2, 3},
	}
	p := ctx.PopPos()
	if p != 3 {
		t.Errorf("PopPos() = %d, want %d", p, 3)
	}
	if len(ctx.positions) != 2 {
		t.Errorf("PopPos() = %d, want %d", len(ctx.positions), 2)
	}
}
