package render

import (
	"fmt"
	"testing"
)

func TestPageInner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &hookBase{}
	c.pageInner = "pageInner"
	if c.PageInner() != "pageInner" {
		t.Errorf("PageInner() = %v, want %v", c.PageInner(), "pageInner")
	}
}
