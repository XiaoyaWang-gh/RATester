package page

import (
	"fmt"
	"testing"
)

func TestType_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expected := ""
	result := nop.Type()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
