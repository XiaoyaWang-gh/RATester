package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestStaleVersion_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &PagesFromTemplate{
		buildState: &BuildState{
			StaleVersion: 1,
		},
	}

	expected := uint32(1)
	actual := b.StaleVersion()

	if actual != expected {
		t.Errorf("Expected StaleVersion to be %d, but got %d", expected, actual)
	}
}
