package logs

import (
	"fmt"
	"testing"
)

func TestnewSLACKWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := newSLACKWriter()
	if logger == nil {
		t.Error("Expected logger to be not nil")
	}
}
