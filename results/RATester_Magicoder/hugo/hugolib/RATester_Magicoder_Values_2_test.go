package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValues_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	taxonomies := taxonomiesConfig{
		"singular1": "plural1",
		"singular2": "plural2",
	}

	expectedViews := []viewName{
		{singular: "singular1", plural: "plural1", pluralTreeKey: "plural1"},
		{singular: "singular2", plural: "plural2", pluralTreeKey: "plural2"},
	}

	expectedViewsByTreeKey := map[string]viewName{
		"plural1": {singular: "singular1", plural: "plural1", pluralTreeKey: "plural1"},
		"plural2": {singular: "singular2", plural: "plural2", pluralTreeKey: "plural2"},
	}

	result := taxonomies.Values()

	if !reflect.DeepEqual(result.views, expectedViews) {
		t.Errorf("Expected views to be %v, but got %v", expectedViews, result.views)
	}

	if !reflect.DeepEqual(result.viewsByTreeKey, expectedViewsByTreeKey) {
		t.Errorf("Expected viewsByTreeKey to be %v, but got %v", expectedViewsByTreeKey, result.viewsByTreeKey)
	}
}
