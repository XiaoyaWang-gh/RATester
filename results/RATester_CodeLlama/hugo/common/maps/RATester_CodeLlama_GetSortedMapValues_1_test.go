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

	c := &Scratch{
		values: map[string]any{
			"key1": map[string]any{
				"key11": "value11",
				"key12": "value12",
				"key13": "value13",
			},
			"key2": map[string]any{
				"key21": "value21",
				"key22": "value22",
				"key23": "value23",
			},
		},
	}

	expected := []any{
		"value11",
		"value12",
		"value13",
		"value21",
		"value22",
		"value23",
	}

	actual := c.GetSortedMapValues("key1")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
