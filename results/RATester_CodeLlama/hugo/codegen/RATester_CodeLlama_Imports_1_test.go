package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestImports_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Methods{
		{
			Imports: []string{"fmt"},
		},
		{
			Imports: []string{"fmt", "os"},
		},
		{
			Imports: []string{"fmt", "os", "io"},
		},
	}

	expected := []string{"fmt", "os", "io"}
	actual := m.Imports()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
