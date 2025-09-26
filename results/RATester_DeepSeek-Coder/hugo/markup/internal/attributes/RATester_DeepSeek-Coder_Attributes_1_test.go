package attributes

import (
	"fmt"
	"testing"
)

func TestAttributes_1(t *testing.T) {
	t.Run("TestAttributes", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		holder := &AttributesHolder{
			attributes: []Attribute{
				{Name: "name1", Value: "value1"},
				{Name: "name2", Value: "value2"},
			},
		}

		attributes := holder.Attributes()
		if len(attributes) != 2 {
			t.Errorf("Expected 2 attributes, got %d", len(attributes))
		}

		if attributes["name1"] != "value1" {
			t.Errorf("Expected 'value1', got '%v'", attributes["name1"])
		}

		if attributes["name2"] != "value2" {
			t.Errorf("Expected 'value2', got '%v'", attributes["name2"])
		}
	})
}
