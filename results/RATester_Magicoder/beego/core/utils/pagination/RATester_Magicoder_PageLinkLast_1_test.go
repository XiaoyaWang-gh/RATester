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

	req, _ := http.NewRequest("GET", "/", nil)
	paginator := &Paginator{
		Request:     req,
		PerPageNums: 10,
		MaxPages:    100,
		nums:        1000,
	}
	paginator.PageNums()
	expected := "/?page=100"
	actual := paginator.PageLinkLast()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
