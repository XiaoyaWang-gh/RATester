package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestReset_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{
		contentRenderedVersion: 1,
		contentRendered:        atomic.Bool{},
		renderHooks:            &renderHooks{},
	}

	pco.Reset()

	if pco.contentRenderedVersion != 2 {
		t.Errorf("Expected contentRenderedVersion to be 2, but got %d", pco.contentRenderedVersion)
	}

	if pco.contentRendered.Load() {
		t.Error("Expected contentRendered to be false, but got true")
	}

	if pco.renderHooks == nil {
		t.Error("Expected renderHooks to not be nil, but got nil")
	}
}
