package render

import (
	"fmt"
	"testing"
)

func TestPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &hookBase{}
	c.page = "page"
	if c.Page() != "page" {
		t.Errorf("Page() = %v, want %v", c.Page(), "page")
	}
}
