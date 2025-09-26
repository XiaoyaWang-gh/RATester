package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestbindMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Input: &BeegoInput{
				Context: &Context{},
			},
		},
	}

	params := &url.Values{}
	params.Add("key[0]", "value0")
	params.Add("key[1]", "value1")

	typ := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(""))
	result := input.bindMap(params, "key", typ)

	expected := map[string]string{
		"0": "value0",
		"1": "value1",
	}

	if !reflect.DeepEqual(result.Interface(), expected) {
		t.Errorf("Expected %v, but got %v", expected, result.Interface())
	}
}
