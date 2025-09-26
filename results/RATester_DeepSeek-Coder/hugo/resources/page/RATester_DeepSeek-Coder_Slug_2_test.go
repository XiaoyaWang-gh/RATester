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
	expected := ""
	actual := p.Slug()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
