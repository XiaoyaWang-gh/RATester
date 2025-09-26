package attributes

import (
	"fmt"
	"testing"
)

func TestOptionsSlice_1(t *testing.T) {
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

	options := holder.OptionsSlice()

	if len(options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(options))
	}

	if options[0].Name != "option1" || options[0].Value != "value1" {
		t.Errorf("Expected option1 with value1, got %s with %s", options[0].Name, options[0].Value)
	}

	if options[1].Name != "option2" || options[1].Value != "value2" {
		t.Errorf("Expected option2 with value2, got %s with %s", options[1].Name, options[1].Value)
	}
}
