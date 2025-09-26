package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    10,
		nums:        100,
		pageRange:   nil,
		pageNums:    10,
		page:        5,
	}
	pages := p.Pages()
	if len(pages) != 9 {
		t.Errorf("Pages() = %v, want %v", len(pages), 9)
	}
	if pages[0] != 1 {
		t.Errorf("Pages()[0] = %v, want %v", pages[0], 1)
	}
	if pages[1] != 2 {
		t.Errorf("Pages()[1] = %v, want %v", pages[1], 2)
	}
	if pages[2] != 3 {
		t.Errorf("Pages()[2] = %v, want %v", pages[2], 3)
	}
	if pages[3] != 4 {
		t.Errorf("Pages()[3] = %v, want %v", pages[3], 4)
	}
	if pages[4] != 5 {
		t.Errorf("Pages()[4] = %v, want %v", pages[4], 5)
	}
	if pages[5] != 6 {
		t.Errorf("Pages()[5] = %v, want %v", pages[5], 6)
	}
	if pages[6] != 7 {
		t.Errorf("Pages()[6] = %v, want %v", pages[6], 7)
	}
	if pages[7] != 8 {
		t.Errorf("Pages()[7] = %v, want %v", pages[7], 8)
	}
	if pages[8] != 9 {
		t.Errorf("Pages()[8] = %v, want %v", pages[8], 9)
	}
}
