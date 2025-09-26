package logs

import (
	"fmt"
	"testing"
)

func TestnewJLWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := newJLWriter()
	if logger == nil {
		t.Error("newJLWriter() should not return nil")
	}
}
