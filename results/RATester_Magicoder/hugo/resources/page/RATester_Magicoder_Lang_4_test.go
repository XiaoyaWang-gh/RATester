package page

import (
	"fmt"
	"testing"
)

func TestLang_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Lang() != "" {
		t.Errorf("Expected Lang() to return an empty string, but got %s", p.Lang())
	}
}
