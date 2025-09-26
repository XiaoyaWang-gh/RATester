package orm

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm/hints"
)

func TestIgnoreIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	indexes := []string{"idx1", "idx2"}
	o.IgnoreIndex(indexes...)
	if o.useIndex != hints.KeyIgnoreIndex {
		t.Errorf("Expected useIndex to be %d, but got %d", hints.KeyIgnoreIndex, o.useIndex)
	}
	if !reflect.DeepEqual(o.indexes, indexes) {
		t.Errorf("Expected indexes to be %v, but got %v", indexes, o.indexes)
	}
}
