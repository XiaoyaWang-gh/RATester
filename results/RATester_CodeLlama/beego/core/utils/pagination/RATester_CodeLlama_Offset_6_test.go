package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOffset_6(t *testing.T) {
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
		page:        1,
	}
	if p.Offset() != 0 {
		t.Errorf("Offset() = %v, want %v", p.Offset(), 0)
	}
}
