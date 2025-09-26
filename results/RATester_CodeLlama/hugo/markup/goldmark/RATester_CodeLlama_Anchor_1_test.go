package goldmark

import (
	"fmt"
	"testing"
)

func TestAnchor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{}
	ctx.anchor = "anchor"
	if ctx.Anchor() != "anchor" {
		t.Errorf("Anchor() = %v, want %v", ctx.Anchor(), "anchor")
	}
}
