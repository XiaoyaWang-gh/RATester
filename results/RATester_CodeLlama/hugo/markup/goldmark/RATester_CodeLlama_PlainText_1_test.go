package goldmark

import (
	"fmt"
	"testing"
)

func TestPlainText_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{}
	ctx.plainText = "plain text"
	if ctx.PlainText() != "plain text" {
		t.Errorf("PlainText() = %v, want %v", ctx.PlainText(), "plain text")
	}
}
