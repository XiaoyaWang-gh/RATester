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

	expected := 10
	actual := paginator.PagerSize()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
