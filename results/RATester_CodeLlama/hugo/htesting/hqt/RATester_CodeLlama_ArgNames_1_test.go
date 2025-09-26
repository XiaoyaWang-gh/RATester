package hqt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := argNames{"a", "b", "c"}
	if !reflect.DeepEqual(a.ArgNames(), []string{"a", "b", "c"}) {
		t.Errorf("ArgNames() = %v, want %v", a.ArgNames(), []string{"a", "b", "c"})
	}
}
