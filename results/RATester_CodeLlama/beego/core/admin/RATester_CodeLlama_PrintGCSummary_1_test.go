package admin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintGCSummary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w bytes.Buffer
	PrintGCSummary(&w)
	if w.String() == "" {
		t.Errorf("PrintGCSummary() = %v, want non-empty string", w.String())
	}
}
