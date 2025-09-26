package parse

import (
	"fmt"
	"testing"
)

func Testexpect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
	}
	expected := item{
		typ:  itemType(1),
		pos:  Pos(1),
		val:  "test",
		line: 1,
	}
	context := "test context"
	tree.expect(expected.typ, context)
}
