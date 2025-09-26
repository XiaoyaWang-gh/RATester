package page

import (
	"fmt"
	"testing"
)

func TestContentBaseName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	result := p.ContentBaseName()
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
