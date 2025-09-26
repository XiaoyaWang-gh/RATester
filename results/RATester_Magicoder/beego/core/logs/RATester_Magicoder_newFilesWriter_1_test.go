package logs

import (
	"fmt"
	"testing"
)

func TestnewFilesWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := newFilesWriter()
	if logger == nil {
		t.Error("Expected a non-nil Logger, got nil")
	}
}
