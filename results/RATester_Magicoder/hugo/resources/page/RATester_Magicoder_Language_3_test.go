package page

import (
	"fmt"
	"testing"
)

func TestLanguage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Language() != nil {
		t.Errorf("Expected Language() to return nil, but got %v", p.Language())
	}
}
