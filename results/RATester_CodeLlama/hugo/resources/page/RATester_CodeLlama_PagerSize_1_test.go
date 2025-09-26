package page

import (
	"fmt"
	"testing"
)

func TestPagerSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	p.size = 10
	if p.PagerSize() != 10 {
		t.Errorf("PagerSize() = %v, want %v", p.PagerSize(), 10)
	}
}
