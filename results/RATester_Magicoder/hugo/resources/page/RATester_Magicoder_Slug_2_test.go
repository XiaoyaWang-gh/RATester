package page

import (
	"fmt"
	"testing"
)

func TestSlug_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Slug() != "" {
		t.Errorf("Expected Slug() to return an empty string, but got %s", p.Slug())
	}
}
