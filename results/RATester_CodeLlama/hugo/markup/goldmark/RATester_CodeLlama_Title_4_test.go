package goldmark

import (
	"fmt"
	"testing"
)

func TestTitle_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := linkContext{}
	ctx.title = "title"
	if ctx.Title() != "title" {
		t.Errorf("Title() = %v, want %v", ctx.Title(), "title")
	}
}
