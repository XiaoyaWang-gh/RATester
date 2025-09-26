package echo

import (
	"fmt"
	"testing"
)

func TestSetParamNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context{
		pnames: []string{"name1", "name2"},
	}

	ctx.SetParamNames("name3", "name4")

	if len(ctx.pnames) != 2 {
		t.Errorf("Expected pnames length to be 2, but got %d", len(ctx.pnames))
	}

	if ctx.pnames[0] != "name3" || ctx.pnames[1] != "name4" {
		t.Errorf("Expected pnames to be [name3, name4], but got %v", ctx.pnames)
	}
}
