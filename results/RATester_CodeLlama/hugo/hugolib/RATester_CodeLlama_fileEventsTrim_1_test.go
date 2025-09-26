package hugolib

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestFileEventsTrim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var h *HugoSites
	var events []fsnotify.Event
	if got := h.fileEventsTrim(events); got != nil {
		t.Errorf("fileEventsTrim() = %v, want nil", got)
	}
}
