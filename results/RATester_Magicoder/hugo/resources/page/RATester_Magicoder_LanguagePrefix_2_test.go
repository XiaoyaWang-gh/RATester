package page

import (
	"fmt"
	"testing"
)

func TestLanguagePrefix_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expected := ""
	result := nop.LanguagePrefix()
	if result != expected {
		t.Errorf("Expected LanguagePrefix to be %s, but got %s", expected, result)
	}
}
