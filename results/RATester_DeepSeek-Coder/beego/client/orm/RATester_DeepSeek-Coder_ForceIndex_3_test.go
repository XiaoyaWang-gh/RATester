package orm

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestForceIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{
		useIndex: hints.KeyForceIndex,
		indexes:  []string{"index1", "index2"},
	}

	newO := o.ForceIndex("index3", "index4")

	if newO.(*querySet).useIndex != hints.KeyForceIndex {
		t.Errorf("Expected useIndex to be %v, got %v", hints.KeyForceIndex, newO.(*querySet).useIndex)
	}

	if len(newO.(*querySet).indexes) != 4 {
		t.Errorf("Expected 4 indexes, got %v", len(newO.(*querySet).indexes))
	}

	if newO.(*querySet).indexes[0] != "index3" || newO.(*querySet).indexes[1] != "index4" {
		t.Errorf("Expected indexes to be ['index3', 'index4'], got %v", newO.(*querySet).indexes)
	}
}
