package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "true"
	typ := reflect.TypeOf(true)
	got := input.bindBool(val, typ)
	want := reflect.ValueOf(true)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input.bindBool() = %v, want %v", got, want)
	}
}
