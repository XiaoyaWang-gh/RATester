package pagination

import (
	"fmt"
	"testing"
)

func TestHasPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		PerPageNums: 10,
		nums:        100,
	}

	hasPages := paginator.HasPages()
	if !hasPages {
		t.Errorf("Expected HasPages to be true, got false")
	}

	paginator.nums = 0
	hasPages = paginator.HasPages()
	if hasPages {
		t.Errorf("Expected HasPages to be false, got true")
	}
}
