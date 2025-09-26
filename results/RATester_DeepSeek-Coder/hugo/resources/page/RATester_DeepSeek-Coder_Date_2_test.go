package page

import (
	"fmt"
	"testing"
)

func TestDate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	nop := nopPage(0)
	date := nop.Date()

	if date.IsZero() {
		t.Errorf("Expected non-zero time, got %v", date)
	}
}
