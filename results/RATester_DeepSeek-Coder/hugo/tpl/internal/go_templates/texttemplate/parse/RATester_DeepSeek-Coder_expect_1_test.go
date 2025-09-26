package parse

import (
	"fmt"
	"testing"
)

func TestExpect_1(t *testing.T) {
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
		pos:  1,
		val:  "test",
		line: 1,
	}

	tree.expect(expected.typ, "test context")

	// Test with unexpected item
	unexpected := item{
		typ:  itemType(2),
		pos:  2,
		val:  "unexpected",
		line: 2,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	tree.expect(unexpected.typ, "test context")
}
