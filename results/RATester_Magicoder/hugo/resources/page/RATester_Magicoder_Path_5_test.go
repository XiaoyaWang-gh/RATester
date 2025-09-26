package page

import (
	"fmt"
	"testing"
)

func TestPath_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expected := ""
	result := nop.Path()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
