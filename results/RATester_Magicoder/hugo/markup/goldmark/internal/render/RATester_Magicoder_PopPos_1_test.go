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

	expected := 3
	actual := ctx.PopPos()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}

	if len(ctx.positions) != 2 {
		t.Errorf("Expected positions length to be 2, but got %d", len(ctx.positions))
	}
}
