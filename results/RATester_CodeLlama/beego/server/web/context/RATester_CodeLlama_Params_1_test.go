package context

import (
	"fmt"
	"testing"
)

func TestParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames:  []string{"name", "age"},
		pvalues: []string{"astaxie", "25"},
	}
	m := input.Params()
	if m["name"] != "astaxie" || m["age"] != "25" {
		t.Error("Params failed")
	}
}
