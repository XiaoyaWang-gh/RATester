package orm

import (
	"fmt"
	"testing"
)

func TestFilterRaw_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	expr := "expr"
	sql := "sql"
	o.FilterRaw(expr, sql)
	if o.cond == nil {
		t.Errorf("TestFilterRaw failed")
	}
}
