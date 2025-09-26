package mock

import (
	"fmt"
	"testing"
)

func TestGroupBy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}

	exprs := []string{"id"}

	result := d.GroupBy(exprs...)

	if result != d {
		t.Errorf("Expected %v, but got %v", d, result)
	}
}
