package codegen

import (
	"fmt"
	"testing"
)

func TestOutStrNamed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Method{
		Out: []string{"string"},
	}
	if got := m.outStrNamed(); got != "(o0 string)" {
		t.Errorf("outStrNamed() = %v, want %v", got, "(o0 string)")
	}
}
