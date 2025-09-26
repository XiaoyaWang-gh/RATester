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

	b := &PagesFromTemplate{}
	if b.StaleVersion() != 0 {
		t.Errorf("StaleVersion() = %d, want 0", b.StaleVersion())
	}
	b.MarkStale()
	if b.StaleVersion() != 1 {
		t.Errorf("StaleVersion() = %d, want 1", b.StaleVersion())
	}
}
