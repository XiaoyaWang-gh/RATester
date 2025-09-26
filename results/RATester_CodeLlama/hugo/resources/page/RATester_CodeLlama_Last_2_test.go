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
	}
	if p.Last() != p {
		t.Errorf("Pager.Last() = %v, want %v", p.Last(), p)
	}
}
