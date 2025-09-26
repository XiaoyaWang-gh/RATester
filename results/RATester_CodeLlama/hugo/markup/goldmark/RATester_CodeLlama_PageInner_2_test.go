package goldmark

import (
	"fmt"
	"testing"
)

func TestPageInner_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{}
	ctx.pageInner = "pageInner"
	if ctx.PageInner() != "pageInner" {
		t.Errorf("PageInner() = %v, want %v", ctx.PageInner(), "pageInner")
	}
}
