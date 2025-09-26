package pagination

import (
	"fmt"
	"testing"
)

func TestPageNums_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{
		PerPageNums: 10,
		MaxPages:    10,
		nums:        100,
	}
	if p.PageNums() != 10 {
		t.Errorf("PageNums() = %v, want %v", p.PageNums(), 10)
	}
}
