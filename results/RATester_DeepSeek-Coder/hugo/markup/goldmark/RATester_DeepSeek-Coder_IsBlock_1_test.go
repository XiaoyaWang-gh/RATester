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

	ctx := imageLinkContext{
		linkContext: linkContext{},
		ordinal:     1,
		isBlock:     true,
	}

	if !ctx.IsBlock() {
		t.Errorf("Expected IsBlock to return true, got false")
	}
}
