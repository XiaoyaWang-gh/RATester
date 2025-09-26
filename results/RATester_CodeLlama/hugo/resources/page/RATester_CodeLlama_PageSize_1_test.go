package page

import (
	"fmt"
	"testing"
)

func TestPageSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	p.size = 10
	if p.PageSize() != 10 {
		t.Errorf("PageSize() = %d, want %d", p.PageSize(), 10)
	}
}
