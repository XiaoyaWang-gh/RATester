package page

import (
	"fmt"
	"testing"
)

func TestIsHome_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsHome() != false {
		t.Errorf("Expected IsHome() to return false, but got true")
	}
}
