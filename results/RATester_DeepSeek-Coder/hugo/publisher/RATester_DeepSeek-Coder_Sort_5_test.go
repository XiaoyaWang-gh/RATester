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
		Tags:    []string{"div", "span", "p"},
		Classes: []string{"class1", "class2", "class3"},
		IDs:     []string{"id1", "id2", "id3"},
	}

	h.Sort()

	expectedTags := []string{"div", "p", "span"}
	expectedClasses := []string{"class1", "class2", "class3"}
	expectedIDs := []string{"id1", "id2", "id3"}

	if !reflect.DeepEqual(h.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, got %v", expectedTags, h.Tags)
	}

	if !reflect.DeepEqual(h.Classes, expectedClasses) {
		t.Errorf("Expected Classes to be %v, got %v", expectedClasses, h.Classes)
	}

	if !reflect.DeepEqual(h.IDs, expectedIDs) {
		t.Errorf("Expected IDs to be %v, got %v", expectedIDs, h.IDs)
	}
}
