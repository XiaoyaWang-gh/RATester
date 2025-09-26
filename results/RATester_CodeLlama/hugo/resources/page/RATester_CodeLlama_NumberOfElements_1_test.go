package page

import (
	"fmt"
	"testing"
)

func TestNumberOfElements_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pager{
		number: 1,
	}
	if got := p.NumberOfElements(); got != 0 {
		t.Errorf("Pager.NumberOfElements() = %v, want %v", got, 0)
	}
}
