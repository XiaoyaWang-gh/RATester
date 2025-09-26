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

	engine := &Engine{}

	left := "<%"
	right := "%>"

	engine.Delims(left, right)

	if engine.delims.Left != left || engine.delims.Right != right {
		t.Errorf("Expected delims to be %s and %s, but got %s and %s", left, right, engine.delims.Left, engine.delims.Right)
	}
}
