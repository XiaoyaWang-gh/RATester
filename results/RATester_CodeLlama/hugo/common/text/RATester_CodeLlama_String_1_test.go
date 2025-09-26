package text

import (
	"fmt"
	"testing"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pos := Position{
		Filename:     "test.go",
		Offset:       10,
		LineNumber:   10,
		ColumnNumber: 10,
	}
	if pos.String() != "test.go:10:10" {
		t.Errorf("pos.String() = %q, want %q", pos.String(), "test.go:10:10")
	}
}
