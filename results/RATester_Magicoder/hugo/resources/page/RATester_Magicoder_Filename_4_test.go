package page

import (
	"fmt"
	"testing"
)

func TestFilename_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	filename := nop.Filename()
	if filename != "" {
		t.Errorf("Expected Filename to be empty, but got %s", filename)
	}
}
