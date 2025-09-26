package gin

import (
	"fmt"
	"testing"
)

func TestDelims_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()

	left := "<%"
	right := "%>"

	engine.Delims(left, right)

	if engine.delims.Left != left || engine.delims.Right != right {
		t.Errorf("Expected delims to be %s%s, but got %s%s", left, right, engine.delims.Left, engine.delims.Right)
	}
}
