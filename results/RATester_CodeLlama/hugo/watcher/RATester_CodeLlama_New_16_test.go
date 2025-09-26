package watcher

import (
	"fmt"
	"testing"
	"time"
)

func TestNew_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	intervalBatcher := time.Duration(10)
	intervalPoll := time.Duration(10)
	poll := true

	// Act
	batcher, err := New(intervalBatcher, intervalPoll, poll)

	// Assert
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	if batcher == nil {
		t.Errorf("New() batcher = nil")
		return
	}

	if batcher.FileWatcher == nil {
		t.Errorf("New() batcher.FileWatcher = nil")
		return
	}

	if batcher.ticker == nil {
		t.Errorf("New() batcher.ticker = nil")
		return
	}

	if batcher.done == nil {
		t.Errorf("New() batcher.done = nil")
		return
	}

	if batcher.Events == nil {
		t.Errorf("New() batcher.Events = nil")
		return
	}
}
