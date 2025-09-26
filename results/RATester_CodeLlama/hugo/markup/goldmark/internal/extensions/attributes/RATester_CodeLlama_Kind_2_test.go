package attributes

import (
	"fmt"
	"testing"
)

func TestKind_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &attributesBlock{}
	if a.Kind() != kindAttributesBlock {
		t.Errorf("Expected %d, but actual %d", kindAttributesBlock, a.Kind())
	}
}
