package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLen_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pairList{
		Pairs: []pair{
			{reflect.ValueOf("a"), reflect.ValueOf(1)},
			{reflect.ValueOf("b"), reflect.ValueOf(2)},
			{reflect.ValueOf("c"), reflect.ValueOf(3)},
		},
	}
	if p.Len() != 3 {
		t.Errorf("Expected 3, got %d", p.Len())
	}
}
