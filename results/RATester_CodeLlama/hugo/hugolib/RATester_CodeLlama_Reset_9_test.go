package hugolib

import (
	"fmt"
	"testing"
)

func TestReset_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	pco.Reset()
	if pco.contentRenderedVersion != 1 {
		t.Errorf("contentRenderedVersion should be 1, got %d", pco.contentRenderedVersion)
	}
	if !pco.contentRendered.Load() {
		t.Error("contentRendered should be true")
	}
	if pco.renderHooks == nil {
		t.Error("renderHooks should not be nil")
	}
}
