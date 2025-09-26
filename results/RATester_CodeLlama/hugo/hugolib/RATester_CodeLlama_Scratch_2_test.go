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
	if scp.Scratch() == nil {
		t.Error("Scratch() should not return nil")
	}
}
