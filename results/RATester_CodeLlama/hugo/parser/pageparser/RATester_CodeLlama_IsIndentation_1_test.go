package pageparser

import (
	"fmt"
	"testing"
)

func TestIsIndentation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := Item{Type: tIndentation}
	if !i.IsIndentation() {
		t.Errorf("IsIndentation failed")
	}
}
