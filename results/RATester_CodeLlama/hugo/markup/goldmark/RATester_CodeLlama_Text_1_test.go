package goldmark

import (
	"fmt"
	"testing"
)

func TestText_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{}
	ctx.text = "text"
	if ctx.Text() != "text" {
		t.Errorf("Text() = %v, want %v", ctx.Text(), "text")
	}
}
