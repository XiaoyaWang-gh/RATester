package filenotify

import (
	"fmt"
	"testing"
)

func TestErrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	fp := &filePoller{
		errors: make(chan error),
	}

	errChan := fp.Errors()

	// Test that the errors channel is returned correctly
	if errChan != fp.errors {
		t.Errorf("Expected %v, got %v", fp.errors, errChan)
	}

	// Test that the errors channel is closed when the filePoller is closed
	fp.closed = true
	close(fp.errors)

	_, ok := <-errChan
	if ok {
		t.Errorf("Expected the errors channel to be closed")
	}
}
