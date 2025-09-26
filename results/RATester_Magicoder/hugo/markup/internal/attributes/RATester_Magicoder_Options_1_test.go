package attributes

import (
	"fmt"
	"testing"
)

func TestOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	holder := &AttributesHolder{
		options: []Attribute{
			{Name: "option1", Value: "value1"},
			{Name: "option2", Value: "value2"},
		},
	}

	options := holder.Options()

	if len(options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(options))
	}

	if options["option1"] != "value1" {
		t.Errorf("Expected option1 to be 'value1', got %v", options["option1"])
	}

	if options["option2"] != "value2" {
		t.Errorf("Expected option2 to be 'value2', got %v", options["option2"])
	}
}
