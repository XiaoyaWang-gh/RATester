package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEndControl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
	}

	expected := tree.newEnd(tree.expect(itemRightDelim, "end").pos)
	result := tree.endControl()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
