package commands

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/hugolib/filesystems"
)

func TestPartitionDynamicEvents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sourceFs := &filesystems.SourceFilesystems{}
	events := []fsnotify.Event{}
	de := partitionDynamicEvents(sourceFs, events)
	if len(de.ContentEvents) != 0 {
		t.Errorf("Expected 0 content events, got %d", len(de.ContentEvents))
	}
	if len(de.AssetEvents) != 0 {
		t.Errorf("Expected 0 asset events, got %d", len(de.AssetEvents))
	}
}
