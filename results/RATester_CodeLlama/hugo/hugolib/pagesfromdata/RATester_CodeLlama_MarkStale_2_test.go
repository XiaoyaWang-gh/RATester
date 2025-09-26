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

	b := &PagesFromTemplate{}
	b.MarkStale()
	if b.buildState.StaleVersion != 1 {
		t.Errorf("StaleVersion should be 1, got %d", b.buildState.StaleVersion)
	}
}
