package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pairList := pairList{
		Pairs: []pair{
			{Key: reflect.ValueOf(1), Value: reflect.ValueOf("one")},
			{Key: reflect.ValueOf(2), Value: reflect.ValueOf("two")},
		},
	}

	pairList.Swap(0, 1)

	if pairList.Pairs[0].Key.Interface() != 2 || pairList.Pairs[0].Value.Interface() != "two" {
		t.Errorf("Expected Pairs[0] to be {2, \"two\"}, but got {%v, %v}", pairList.Pairs[0].Key.Interface(), pairList.Pairs[0].Value.Interface())
	}

	if pairList.Pairs[1].Key.Interface() != 1 || pairList.Pairs[1].Value.Interface() != "one" {
		t.Errorf("Expected Pairs[1] to be {1, \"one\"}, but got {%v, %v}", pairList.Pairs[1].Key.Interface(), pairList.Pairs[1].Value.Interface())
	}
}
