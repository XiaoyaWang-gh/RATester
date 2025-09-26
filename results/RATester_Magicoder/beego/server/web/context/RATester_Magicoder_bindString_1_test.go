package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbindString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "test"
	typ := reflect.TypeOf(val)
	result := input.bindString(val, typ)
	if result.String() != val {
		t.Errorf("Expected %s, got %s", val, result.String())
	}
}
