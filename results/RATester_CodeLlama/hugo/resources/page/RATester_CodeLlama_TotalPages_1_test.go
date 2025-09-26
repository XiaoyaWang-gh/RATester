package page

import (
	"fmt"
	"testing"
)

func TestTotalPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	p.paginatedElements = []paginatedElement{}
	if got := p.TotalPages(); got != 0 {
		t.Errorf("TotalPages() = %v, want %v", got, 0)
	}
}
