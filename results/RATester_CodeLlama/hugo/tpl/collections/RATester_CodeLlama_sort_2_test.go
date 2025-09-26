package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort_2(t *testing.T) {
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
			{reflect.ValueOf("d"), reflect.ValueOf(4)},
			{reflect.ValueOf("e"), reflect.ValueOf(5)},
		},
		SortAsc: true,
	}
	sorted := p.sort()
	if !reflect.DeepEqual(sorted, []interface{}{1, 2, 3, 4, 5}) {
		t.Errorf("Expected %v, got %v", []interface{}{1, 2, 3, 4, 5}, sorted)
	}
}
