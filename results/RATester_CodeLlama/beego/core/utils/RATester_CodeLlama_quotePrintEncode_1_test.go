package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestQuotePrintEncode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var buf bytes.Buffer
	err := quotePrintEncode(&buf, "Hello, World!")
	if err != nil {
		t.Errorf("quotePrintEncode() error = %v", err)
		return
	}
	if buf.String() != "Hello, World!\r\n" {
		t.Errorf("quotePrintEncode() = %v, want %v", buf.String(), "Hello, World!\r\n")
	}
}
