package page

import (
	"fmt"
	"testing"
)

func TestDir_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	expected := ""
	if result := p.Dir(); result != expected {
		t.Errorf("Expected Dir() to return %s, but got %s", expected, result)
	}
}
