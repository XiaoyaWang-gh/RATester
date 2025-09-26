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

	elements := HTMLElements{
		Tags:    []string{"a", "c", "b"},
		Classes: []string{"d", "e", "f"},
		IDs:     []string{"g", "h", "i"},
	}

	elements.Sort()

	expected := HTMLElements{
		Tags:    []string{"a", "b", "c"},
		Classes: []string{"d", "e", "f"},
		IDs:     []string{"g", "h", "i"},
	}

	if !reflect.DeepEqual(elements, expected) {
		t.Errorf("Expected %v, got %v", expected, elements)
	}
}
