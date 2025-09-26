package orm

import (
	"fmt"
	"testing"
)

func TestUseIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := &querySet{}
	indexes := []string{"index1", "index2"}
	qs.UseIndex(indexes...)

	if len(qs.indexes) != len(indexes) {
		t.Errorf("Expected %d indexes, got %d", len(indexes), len(qs.indexes))
	}

	for i, index := range indexes {
		if qs.indexes[i] != index {
			t.Errorf("Expected index %s, got %s", index, qs.indexes[i])
		}
	}
}
