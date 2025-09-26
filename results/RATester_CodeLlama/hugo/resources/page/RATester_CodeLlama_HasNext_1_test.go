package page

import (
	"fmt"
	"testing"
)

func TestHasNext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pager{
		number: 1,
	}
	if p.HasNext() {
		t.Errorf("Pager.HasNext() = true, want false")
	}
}
