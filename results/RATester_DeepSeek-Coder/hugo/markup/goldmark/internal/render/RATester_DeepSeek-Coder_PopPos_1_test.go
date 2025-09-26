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
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	expected = 2
	actual = ctx.positions[len(ctx.positions)-1]

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
