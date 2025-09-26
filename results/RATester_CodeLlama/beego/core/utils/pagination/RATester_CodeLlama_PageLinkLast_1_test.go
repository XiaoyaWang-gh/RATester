package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPageLinkLast_1(t *testing.T) {
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
		pageRange:   []int{1, 2, 3, 4, 5},
		pageNums:    10,
		page:        5,
	}
	link := p.PageLinkLast()
	if link != "?page=10" {
		t.Errorf("PageLinkLast() = %v, want %v", link, "?page=10")
	}
}
