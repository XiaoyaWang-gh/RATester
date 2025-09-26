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

	ctx := headingContext{
		page:      "testPage",
		pageInner: "testPageInner",
		level:     1,
		anchor:    "testAnchor",
		text:      "testText",
		plainText: "testPlainText",
	}

	result := ctx.PageInner()

	if result != "testPageInner" {
		t.Errorf("Expected PageInner to return 'testPageInner', but got %v", result)
	}
}
