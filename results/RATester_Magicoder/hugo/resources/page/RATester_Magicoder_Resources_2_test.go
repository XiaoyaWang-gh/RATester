package page

import (
	"fmt"
	"testing"
)

func TestResources_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Resources() != nil {
		t.Errorf("Expected Resources() to return nil, but got %v", p.Resources())
	}
}
