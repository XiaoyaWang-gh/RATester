package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContinueControl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		rangeDepth: 1,
	}

	pos := Pos(1)
	line := 1

	expected := tree.newContinue(pos, line)
	result := tree.continueControl(pos, line)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
