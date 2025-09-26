package mock

import (
	"fmt"
	"testing"
)

func TestOrderBy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var d *DoNothingQuerySetter
	var exprs []string
	d.OrderBy(exprs...)
	if got := d.OrderBy(exprs...); got != d {
		t.Errorf("DoNothingQuerySetter.OrderBy() = %v, want %v", got, d)
	}
}
