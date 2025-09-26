package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAttributesSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &AttributesHolder{
		attributes: []Attribute{
			{Name: "a", Value: "1"},
			{Name: "b", Value: "2"},
			{Name: "c", Value: "3"},
		},
	}

	if got := a.AttributesSlice(); !reflect.DeepEqual(got, a.attributes) {
		t.Errorf("AttributesSlice() = %v, want %v", got, a.attributes)
	}
}
