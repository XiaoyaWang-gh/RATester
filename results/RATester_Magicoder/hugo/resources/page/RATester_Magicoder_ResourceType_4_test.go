package page

import (
	"fmt"
	"testing"
)

func TestResourceType_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	expected := ""
	if result := p.ResourceType(); result != expected {
		t.Errorf("Expected ResourceType to be %s, but got %s", expected, result)
	}
}
