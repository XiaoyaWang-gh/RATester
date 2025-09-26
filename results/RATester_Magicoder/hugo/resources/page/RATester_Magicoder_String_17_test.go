package page

import (
	"fmt"
	"testing"
)

func TestString_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expected := "nopPage"
	result := nop.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
