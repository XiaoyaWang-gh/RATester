package template

import (
	"fmt"
	"testing"
)

func TestAsciiAlphaNum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c byte
	if asciiAlphaNum(c) {
		t.Errorf("asciiAlphaNum(%c) = true, want false", c)
	}
}
