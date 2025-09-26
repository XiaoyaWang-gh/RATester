package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	val := "123.456"
	typ := reflect.TypeOf(123.456)
	pValue := input.bindFloat(val, typ)
	if pValue.Float() != 123.456 {
		t.Errorf("bindFloat failed, got %v, want %v", pValue.Float(), 123.456)
	}
}
