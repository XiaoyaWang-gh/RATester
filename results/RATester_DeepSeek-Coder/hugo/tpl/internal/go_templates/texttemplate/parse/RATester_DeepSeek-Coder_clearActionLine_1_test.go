package parse

import (
	"fmt"
	"testing"
)

func TestClearActionLine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		actionLine: 10,
	}

	tree.clearActionLine()

	if tree.actionLine != 0 {
		t.Errorf("Expected actionLine to be 0, got %d", tree.actionLine)
	}
}
