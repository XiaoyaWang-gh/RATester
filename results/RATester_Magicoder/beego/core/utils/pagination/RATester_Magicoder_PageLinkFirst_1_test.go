package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPageLinkFirst_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		Request:     &http.Request{},
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
		pageRange:   []int{1, 2, 3},
		pageNums:    10,
		page:        5,
	}

	expected := "/?page=1"
	actual := paginator.PageLinkFirst()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
