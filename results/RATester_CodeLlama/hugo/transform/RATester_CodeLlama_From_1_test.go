package transform

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ft := fromToBuffer{
		from: bytes.NewBufferString("from"),
		to:   bytes.NewBufferString("to"),
	}
	if ft.From() != ft.from {
		t.Errorf("From() = %v, want %v", ft.From(), ft.from)
	}
}
