package page

import (
	"fmt"
	"testing"
)

func TestExtension_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Extension() != "" {
		t.Errorf("Expected Extension() to return an empty string, but got %s", p.Extension())
	}
}
