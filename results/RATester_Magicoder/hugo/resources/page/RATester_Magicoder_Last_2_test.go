package page

import (
	"fmt"
	"testing"
)

func TestLast_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pager{
		number: 1,
		Paginator: &Paginator{
			pagers: []*Pager{
				{number: 1},
				{number: 2},
				{number: 3},
			},
		},
	}

	lastPager := p.Last()
	if lastPager.number != 3 {
		t.Errorf("Expected last pager number to be 3, but got %d", lastPager.number)
	}
}
