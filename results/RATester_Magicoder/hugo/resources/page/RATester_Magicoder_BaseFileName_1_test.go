package page

import (
	"fmt"
	"testing"
)

func TestBaseFileName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.BaseFileName() != "" {
		t.Errorf("Expected BaseFileName to return an empty string, but got %s", p.BaseFileName())
	}
}
