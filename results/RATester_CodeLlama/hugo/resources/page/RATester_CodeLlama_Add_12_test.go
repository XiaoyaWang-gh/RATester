package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.Add("a", "b", "c")
	if !reflect.DeepEqual(p.els, []string{"a", "b", "c"}) {
		t.Errorf("els = %v, want %v", p.els, []string{"a", "b", "c"})
	}
}
