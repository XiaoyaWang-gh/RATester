package navigation

import (
	"fmt"
	"testing"
)

func TestLen_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &menuSorter{
		menu: Menu{
			&MenuEntry{},
			&MenuEntry{},
			&MenuEntry{},
		},
	}

	expected := 3
	actual := ms.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
