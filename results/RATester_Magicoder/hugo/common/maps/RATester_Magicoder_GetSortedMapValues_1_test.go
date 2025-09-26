package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSortedMapValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: map[string]any{
			"testMap": map[string]any{
				"b": 2,
				"a": 1,
				"c": 3,
			},
		},
	}

	expected := []any{1, 2, 3}
	result := scratch.GetSortedMapValues("testMap")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
