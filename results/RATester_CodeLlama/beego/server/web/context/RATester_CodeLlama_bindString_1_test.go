package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "test"
	typ := reflect.TypeOf(val)
	reflect.ValueOf(input).MethodByName("bindString").Call([]reflect.Value{reflect.ValueOf(val), reflect.ValueOf(typ)})
	if input.bindString(val, typ) != reflect.ValueOf(val) {
		t.Errorf("bindString() = %v, want %v", input.bindString(val, typ), reflect.ValueOf(val))
	}
}
