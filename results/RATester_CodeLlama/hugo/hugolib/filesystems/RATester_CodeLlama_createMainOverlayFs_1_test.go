package filesystems

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/paths"
)

func TestCreateMainOverlayFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var b *sourceFilesystemsBuilder
	var p *paths.Paths
	var collector *filesystemsCollector
	var err error

	// CONTEXT_0
	b = &sourceFilesystemsBuilder{}
	p = &paths.Paths{}

	// CONTEXT_1
	collector, err = b.createMainOverlayFs(p)
	if err != nil {
		t.Fatalf("Failed to create main overlay fs: %s", err)
	}

	// CONTEXT_2
	if collector.overlayMounts == nil {
		t.Fatal("Expected overlayMounts to be set")
	}

	// CONTEXT_3
	if collector.overlayMountsContent == nil {
		t.Fatal("Expected overlayMountsContent to be set")
	}

	// CONTEXT_4
	if collector.overlayMountsStatic == nil {
		t.Fatal("Expected overlayMountsStatic to be set")
	}

	// CONTEXT_5
	if collector.overlayFull == nil {
		t.Fatal("Expected overlayFull to be set")
	}

	// CONTEXT_6
	if collector.overlayResources == nil {
		t.Fatal("Expected overlayResources to be set")
	}

	// CONTEXT_7
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}
