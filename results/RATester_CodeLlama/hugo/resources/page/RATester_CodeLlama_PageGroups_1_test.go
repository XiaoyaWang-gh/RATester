package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPageGroups_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pager{
		number: 1,
	}

	if got := p.PageGroups(); !reflect.DeepEqual(got, paginatorEmptyPageGroups) {
		t.Errorf("Pager.PageGroups() = %v, want %v", got, paginatorEmptyPageGroups)
	}
}
