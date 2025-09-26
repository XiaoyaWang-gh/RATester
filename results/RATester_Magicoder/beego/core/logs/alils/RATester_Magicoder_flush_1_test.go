package alils

import (
	"fmt"
	"testing"
)

func Testflush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	store := &LogStore{}
	lg := &LogGroup{}
	writer := &aliLSWriter{
		store: store,
	}

	writer.flush(lg)

	// Assert that the LogGroup's Logs slice is empty
	if len(lg.Logs) != 0 {
		t.Errorf("Expected Logs slice to be empty after flush, but it was not.")
	}
}
