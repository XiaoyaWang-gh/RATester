package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "1"
	typ := reflect.TypeOf(1)
	rv := input.bindValue(val, typ)
	if rv.Int() != 1 {
		t.Errorf("TestbindValue failed")
	}
}
