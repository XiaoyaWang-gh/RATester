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

	p := pairList{
		Pairs: []pair{
			{reflect.ValueOf(1), reflect.ValueOf(2)},
			{reflect.ValueOf(3), reflect.ValueOf(4)},
			{reflect.ValueOf(5), reflect.ValueOf(6)},
		},
	}
	p.Swap(0, 1)
	if p.Pairs[0].Value.Interface() != 4 {
		t.Errorf("Expected 4, got %v", p.Pairs[0].Value.Interface())
	}
	if p.Pairs[1].Value.Interface() != 2 {
		t.Errorf("Expected 2, got %v", p.Pairs[1].Value.Interface())
	}
}
