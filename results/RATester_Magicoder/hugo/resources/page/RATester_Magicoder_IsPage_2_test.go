package page

import (
	"fmt"
	"testing"
)

func TestIsPage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsPage() != false {
		t.Errorf("Expected IsPage() to return false, but it didn't")
	}
}
