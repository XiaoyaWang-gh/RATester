package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "123"
	typ := reflect.TypeOf(1)
	pValue := input.bindInt(val, typ)
	if pValue.Int() != 123 {
		t.Errorf("bindInt failed, got %d, want %d", pValue.Int(), 123)
	}
}
