package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPagers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		pagers: []*Pager{
			{number: 1},
			{number: 2},
			{number: 3},
		},
	}

	expected := []*Pager{
		{number: 1},
		{number: 2},
		{number: 3},
	}

	result := paginator.Pagers()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
