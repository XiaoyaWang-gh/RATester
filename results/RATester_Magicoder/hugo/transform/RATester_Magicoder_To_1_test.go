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
		from: bytes.NewBufferString("test"),
		to:   bytes.NewBuffer(nil),
	}

	w := ft.To()
	if w == nil {
		t.Error("To() should not return nil")
	}
}
