package herrors

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintStackTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := bytes.NewBuffer(nil)
	PrintStackTrace(buf)
	if buf.Len() == 0 {
		t.Error("buf.Len() == 0")
	}
}
