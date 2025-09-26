package orm

import (
	"fmt"
	"testing"
)

func TestOrderBy_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var expressions []string
	o.OrderBy(expressions...)
	if o.orders != nil {
		t.Errorf("TestOrderBy failed")
	}
}
