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

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	p := &Paginator{
		Request:     req,
		PerPageNums: 10,
		MaxPages:    10,
		nums:        100,
		pageRange:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		pageNums:    10,
		page:        1,
	}

	expected := "/?page=10"
	actual := p.PageLinkLast()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
