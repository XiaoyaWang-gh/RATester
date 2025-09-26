package page

import (
	"fmt"
	"testing"
)

func TestPaginator_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	_, err := nop.Paginator()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
