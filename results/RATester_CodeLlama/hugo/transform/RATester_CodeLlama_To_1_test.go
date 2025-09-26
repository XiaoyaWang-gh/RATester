package transform

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ft := fromToBuffer{
		from: &bytes.Buffer{},
		to:   &bytes.Buffer{},
	}
	if got := ft.To(); got != ft.to {
		t.Errorf("To() = %v, want %v", got, ft.to)
	}
}
