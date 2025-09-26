package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{
		values: map[string]any{
			"key1": "value1",
			"key2": "value2",
		},
	}

	expected := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	result := scratch.Values()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
