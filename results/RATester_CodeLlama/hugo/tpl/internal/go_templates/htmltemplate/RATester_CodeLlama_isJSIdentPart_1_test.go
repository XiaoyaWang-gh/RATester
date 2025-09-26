package template

import (
	"fmt"
	"testing"
)

func TestIsJSIdentPart_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var r rune
	if isJSIdentPart(r) {
		t.Errorf("isJSIdentPart(%q) = true, want false", r)
	}
}
