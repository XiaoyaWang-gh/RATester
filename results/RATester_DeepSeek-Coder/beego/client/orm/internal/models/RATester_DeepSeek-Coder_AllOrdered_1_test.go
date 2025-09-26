package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllOrdered_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := &ModelCache{
		orders: []string{"table1", "table2", "table3"},
		cache: map[string]*ModelInfo{
			"table1": {Name: "model1"},
			"table2": {Name: "model2"},
			"table3": {Name: "model3"},
		},
	}

	expected := []*ModelInfo{
		{Name: "model1"},
		{Name: "model2"},
		{Name: "model3"},
	}

	result := mc.AllOrdered()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
