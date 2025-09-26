package pagination

import (
	"fmt"
	"testing"
)

func TestHasPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	p.SetNums(100)
	if !p.HasPages() {
		t.Errorf("HasPages() = %v, want %v", p.HasPages(), true)
	}
}
