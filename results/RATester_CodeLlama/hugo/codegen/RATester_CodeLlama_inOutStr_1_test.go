package codegen

import (
	"fmt"
	"testing"
)

func TestInOutStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		In: []string{"arg0", "arg1"},
	}
	if got := m.inOutStr(); got != "("+"arg0, arg1"+")" {
		t.Errorf("inOutStr() = %v, want %v", got, "("+"arg0, arg1"+")")
	}
}
