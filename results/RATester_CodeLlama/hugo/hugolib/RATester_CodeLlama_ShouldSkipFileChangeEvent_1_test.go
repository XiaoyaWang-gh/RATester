package hugolib

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestShouldSkipFileChangeEvent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var h HugoSites
	var ev fsnotify.Event
	var skip bool

	// CONTEXT_0
	h.skipRebuildForFilenames = map[string]bool{
		"a.txt": true,
	}

	// CONTEXT_1
	ev.Name = "a.txt"
	ev.Op = fsnotify.Write
	skip = h.ShouldSkipFileChangeEvent(ev)
	if !skip {
		t.Errorf("expected skip")
	}

	// CONTEXT_2
	ev.Name = "b.txt"
	ev.Op = fsnotify.Write
	skip = h.ShouldSkipFileChangeEvent(ev)
	if skip {
		t.Errorf("expected not skip")
	}

	// CONTEXT_3
	ev.Name = "a.txt"
	ev.Op = fsnotify.Remove
	skip = h.ShouldSkipFileChangeEvent(ev)
	if skip {
		t.Errorf("expected not skip")
	}

	// CONTEXT_4
	ev.Name = "a.txt"
	ev.Op = fsnotify.Rename
	skip = h.ShouldSkipFileChangeEvent(ev)
	if skip {
		t.Errorf("expected not skip")
	}
}
