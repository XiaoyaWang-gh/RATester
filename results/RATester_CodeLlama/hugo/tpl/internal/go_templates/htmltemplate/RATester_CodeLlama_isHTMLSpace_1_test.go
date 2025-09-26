package template

import (
	"fmt"
	"testing"
)

func TestIsHTMLSpace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c byte
	if isHTMLSpace(c) {
		t.Errorf("isHTMLSpace(%d) = true, want false", c)
	}
}
