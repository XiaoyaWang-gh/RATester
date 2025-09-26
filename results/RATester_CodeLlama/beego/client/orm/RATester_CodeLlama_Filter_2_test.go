package orm

import (
	"fmt"
	"testing"
)

func TestFilter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var expr string
	var args []interface{}
	o.Filter(expr, args...)
	if o.cond == nil {
		t.Errorf("TestFilter failed")
	}
}
