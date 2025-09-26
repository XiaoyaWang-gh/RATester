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

	paginator := &Paginator{
		size: 10,
	}

	size := paginator.PagerSize()

	if size != 10 {
		t.Errorf("Expected size to be 10, but got %d", size)
	}
}
