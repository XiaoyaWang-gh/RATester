package goldmark

import (
	"fmt"
	"testing"
)

func TestIsBlock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := imageLinkContext{}
	ctx.isBlock = true
	if !ctx.IsBlock() {
		t.Errorf("IsBlock() = %v, want %v", ctx.IsBlock(), true)
	}
}
