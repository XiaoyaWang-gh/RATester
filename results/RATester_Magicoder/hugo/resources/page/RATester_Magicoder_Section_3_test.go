package page

import (
	"fmt"
	"testing"
)

func TestSection_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Section() != "" {
		t.Errorf("Expected Section() to return an empty string, but got %s", p.Section())
	}
}
