package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLess_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pairList{
		Pairs: []pair{
			{reflect.ValueOf("a"), reflect.ValueOf("b")},
			{reflect.ValueOf("c"), reflect.ValueOf("d")},
			{reflect.ValueOf("e"), reflect.ValueOf("f")},
		},
		SortAsc: true,
	}

	p.sort()

	if !p.Less(0, 1) {
		t.Errorf("Expected Less(0, 1) to be true")
	}

	if !p.Less(1, 2) {
		t.Errorf("Expected Less(1, 2) to be true")
	}

	if p.Less(0, 2) {
		t.Errorf("Expected Less(0, 2) to be false")
	}

	if p.Less(2, 0) {
		t.Errorf("Expected Less(2, 0) to be false")
	}

	if p.Less(2, 1) {
		t.Errorf("Expected Less(2, 1) to be false")
	}
}
