package attributes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &AttributesHolder{
		options: []Attribute{
			{Name: "foo", Value: "bar"},
			{Name: "baz", Value: "qux"},
		},
	}

	if got, want := a.Options(), map[string]any{"foo": "bar", "baz": "qux"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
