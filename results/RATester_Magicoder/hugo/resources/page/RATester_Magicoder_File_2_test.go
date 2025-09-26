package page

import (
	"fmt"
	"testing"
)

func TestFile_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.File() != nil {
		t.Errorf("Expected File() to return nil, but got %v", p.File())
	}
}
