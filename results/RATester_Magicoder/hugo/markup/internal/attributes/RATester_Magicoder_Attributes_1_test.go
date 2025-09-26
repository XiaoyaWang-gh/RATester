package attributes

import (
	"fmt"
	"testing"
)

func TestAttributes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	holder := &AttributesHolder{
		attributes: []Attribute{
			{Name: "test1", Value: "value1"},
			{Name: "test2", Value: "value2"},
		},
	}

	attributes := holder.Attributes()

	if len(attributes) != 2 {
		t.Errorf("Expected 2 attributes, got %d", len(attributes))
	}

	if attributes["test1"] != "value1" {
		t.Errorf("Expected value1 for test1, got %v", attributes["test1"])
	}

	if attributes["test2"] != "value2" {
		t.Errorf("Expected value2 for test2, got %v", attributes["test2"])
	}
}
