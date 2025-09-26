package page

import (
	"fmt"
	"testing"
)

func TestData_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Data() != nil {
		t.Errorf("Expected Data() to return nil, but got %v", p.Data())
	}
}
