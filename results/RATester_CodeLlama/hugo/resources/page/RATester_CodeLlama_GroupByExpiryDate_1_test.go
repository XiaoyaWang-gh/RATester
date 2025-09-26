package page

import (
	"fmt"
	"testing"
)

func TestGroupByExpiryDate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var p Pages
	var format string
	var order []string

	// when
	_, err := p.GroupByExpiryDate(format, order...)

	// then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
