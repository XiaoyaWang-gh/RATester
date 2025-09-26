package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestForceIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &querySet{}
	indexes := []string{"idx_name1", "idx_name2"}
	o.ForceIndex(indexes...)

	if o.useIndex != hints.KeyForceIndex {
		t.Errorf("Expected useIndex to be %d, but got %d", hints.KeyForceIndex, o.useIndex)
	}

	if !reflect.DeepEqual(o.indexes, indexes) {
		t.Errorf("Expected indexes to be %v, but got %v", indexes, o.indexes)
	}
}
