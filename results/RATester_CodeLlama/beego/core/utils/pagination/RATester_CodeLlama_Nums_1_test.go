package pagination

import (
	"fmt"
	"testing"
)

func TestNums_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{}
	p.nums = 100
	if p.Nums() != 100 {
		t.Errorf("Nums() = %v, want %v", p.Nums(), 100)
	}
}
