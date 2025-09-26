package attributes

import (
	"fmt"
	"testing"
)

func TestAttributesSlice_1(t *testing.T) {
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

	slice := holder.AttributesSlice()

	if len(slice) != len(holder.attributes) {
		t.Errorf("Expected slice length to be %d, but got %d", len(holder.attributes), len(slice))
	}

	for i, attr := range slice {
		if attr.Name != holder.attributes[i].Name || attr.Value != holder.attributes[i].Value {
			t.Errorf("Expected attribute %d to be %v, but got %v", i, holder.attributes[i], attr)
		}
	}
}
