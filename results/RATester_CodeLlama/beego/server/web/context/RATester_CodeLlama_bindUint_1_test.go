package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindUint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "123"
	typ := reflect.TypeOf(uint(0))
	result := input.bindUint(val, typ)
	if result.Kind() != reflect.Uint {
		t.Errorf("bindUint failed")
	}
}
