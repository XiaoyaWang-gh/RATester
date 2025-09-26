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
			{reflect.ValueOf(1), reflect.ValueOf(2)},
			{reflect.ValueOf(3), reflect.ValueOf(4)},
			{reflect.ValueOf(5), reflect.ValueOf(6)},
		},
	}

	expected := 3
	actual := p.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
