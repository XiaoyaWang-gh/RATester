package codegen

import (
	"fmt"
	"testing"
)

func TestInStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		In: []string{"string", "int", "interface{}"},
	}
	if got := m.inStr(); got != "(\"string\", int, interface{})" {
		t.Errorf("inStr() = %v, want \"(\"string\", int, interface{})\"", got)
	}
}
