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
		isBlock: true,
	}

	if !ctx.IsBlock() {
		t.Error("Expected IsBlock to return true, but it didn't")
	}
}
