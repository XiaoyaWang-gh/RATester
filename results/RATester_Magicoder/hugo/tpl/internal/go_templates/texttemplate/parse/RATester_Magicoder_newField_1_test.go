package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewField_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	ident := ".x.y"
	expectedIdent := []string{"x", "y"}

	field := tree.newField(pos, ident)

	if field.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, field.Pos)
	}

	if !reflect.DeepEqual(field.Ident, expectedIdent) {
		t.Errorf("Expected Ident to be %v, but got %v", expectedIdent, field.Ident)
	}
}
