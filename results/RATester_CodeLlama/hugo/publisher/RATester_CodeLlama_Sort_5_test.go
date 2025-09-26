package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HTMLElements{
		Tags:    []string{"b", "i", "u"},
		Classes: []string{"c1", "c2", "c3"},
		IDs:     []string{"id1", "id2", "id3"},
	}
	h.Sort()
	if !reflect.DeepEqual(h.Tags, []string{"b", "i", "u"}) {
		t.Errorf("Tags not sorted")
	}
	if !reflect.DeepEqual(h.Classes, []string{"c1", "c2", "c3"}) {
		t.Errorf("Classes not sorted")
	}
	if !reflect.DeepEqual(h.IDs, []string{"id1", "id2", "id3"}) {
		t.Errorf("IDs not sorted")
	}
}
