package hugolib

import (
	"fmt"
	"testing"
)

func TestScratch_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scp := &ShortcodeWithPage{}
	scratch := scp.Scratch()
	if scratch == nil {
		t.Error("Expected scratch not to be nil")
	}
}
