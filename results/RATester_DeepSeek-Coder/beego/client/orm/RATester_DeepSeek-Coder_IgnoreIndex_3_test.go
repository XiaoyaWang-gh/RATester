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

	qs := &querySet{
		useIndex: hints.KeyIgnoreIndex,
		indexes:  []string{"index1", "index2"},
	}

	newQs := qs.IgnoreIndex("index3", "index4")

	if newQs.(*querySet).useIndex != hints.KeyIgnoreIndex {
		t.Errorf("Expected useIndex to be %v, got %v", hints.KeyIgnoreIndex, newQs.(*querySet).useIndex)
	}

	expectedIndexes := []string{"index1", "index2", "index3", "index4"}
	if !reflect.DeepEqual(newQs.(*querySet).indexes, expectedIndexes) {
		t.Errorf("Expected indexes to be %v, got %v", expectedIndexes, newQs.(*querySet).indexes)
	}
}
