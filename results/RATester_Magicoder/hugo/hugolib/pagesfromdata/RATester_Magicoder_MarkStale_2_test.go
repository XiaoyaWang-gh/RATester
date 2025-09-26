package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestMarkStale_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &PagesFromTemplate{
		buildState: &BuildState{
			StaleVersion: 0,
		},
	}

	b.MarkStale()

	if b.buildState.StaleVersion != 1 {
		t.Errorf("Expected StaleVersion to be 1, but got %d", b.buildState.StaleVersion)
	}
}
