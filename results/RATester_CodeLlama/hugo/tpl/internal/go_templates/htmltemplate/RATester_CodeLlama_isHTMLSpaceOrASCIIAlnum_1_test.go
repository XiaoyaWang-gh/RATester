package template

import (
	"fmt"
	"testing"
)

func TestIsHTMLSpaceOrASCIIAlnum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c byte
	if isHTMLSpaceOrASCIIAlnum(c) {
		t.Errorf("isHTMLSpaceOrASCIIAlnum(%d) = true, want false", c)
	}
}
