package orm

import (
	"fmt"
	"testing"
)

func TestGroupBy_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	exprs := []string{"id"}
	qs.GroupBy(exprs...)
	if len(qs.groups) != 1 {
		t.Errorf("TestGroupBy failed")
	}
}
