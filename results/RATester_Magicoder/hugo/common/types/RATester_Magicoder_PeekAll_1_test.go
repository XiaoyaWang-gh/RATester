package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPeekAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &EvictingStringQueue{
		size: 3,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	expected := []string{"c", "b", "a"}
	actual := q.PeekAll()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
