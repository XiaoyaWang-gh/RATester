package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestIgnoreIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	qs.IgnoreIndex("index1", "index2")
	if qs.useIndex != hints.KeyIgnoreIndex {
		t.Errorf("TestIgnoreIndex failed")
	}
	if len(qs.indexes) != 2 {
		t.Errorf("TestIgnoreIndex failed")
	}
	if qs.indexes[0] != "index1" {
		t.Errorf("TestIgnoreIndex failed")
	}
	if qs.indexes[1] != "index2" {
		t.Errorf("TestIgnoreIndex failed")
	}
}
