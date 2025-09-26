package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAttributes_1(t *testing.T) {
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

	if got, want := a.Attributes(), map[string]any{"a": "1", "b": "2", "c": "3"}; !reflect.DeepEqual(got, want) {
		t.Errorf("Attributes() = %v, want %v", got, want)
	}
}
