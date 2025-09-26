package orm

import (
	"fmt"
	"testing"
)

func TestExclude_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	o.Exclude("expr", "args")
	if o.cond == nil {
		t.Errorf("Exclude failed")
	}
}
