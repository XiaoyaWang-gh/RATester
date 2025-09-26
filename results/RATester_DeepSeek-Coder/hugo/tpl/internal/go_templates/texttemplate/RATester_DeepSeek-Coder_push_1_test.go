package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &state{
		vars: make([]variable, 0),
	}

	name := "testVar"
	value := reflect.ValueOf(42)

	s.push(name, value)

	if len(s.vars) != 1 {
		t.Errorf("Expected length of vars to be 1, got %d", len(s.vars))
	}

	if s.vars[0].name != name {
		t.Errorf("Expected name to be %s, got %s", name, s.vars[0].name)
	}

	if !s.vars[0].value.CanInt() || s.vars[0].value.Int() != 42 {
		t.Errorf("Expected value to be 42, got %v", s.vars[0].value)
	}
}
