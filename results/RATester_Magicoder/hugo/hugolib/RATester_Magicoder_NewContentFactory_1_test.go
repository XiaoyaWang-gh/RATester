package hugolib

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewContentFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	cf := NewContentFactory(h)

	if cf.h != h {
		t.Errorf("Expected cf.h to be %v, but got %v", h, cf.h)
	}

	expectedShortcodeReplacerPre := strings.NewReplacer(
		"{{<", "{x{<",
		"{{%", "{x{%",
		">}}", ">}x}",
		"%}}", "%}x}")
	if cf.shortcodeReplacerPre != expectedShortcodeReplacerPre {
		t.Errorf("Expected cf.shortcodeReplacerPre to be %v, but got %v", expectedShortcodeReplacerPre, cf.shortcodeReplacerPre)
	}

	expectedShortcodeReplacerPost := strings.NewReplacer(
		"{x{<", "{{<",
		"{x{%", "{{%",
		">}x}", ">}}",
		"%}x}", "%}}")
	if cf.shortcodeReplacerPost != expectedShortcodeReplacerPost {
		t.Errorf("Expected cf.shortcodeReplacerPost to be %v, but got %v", expectedShortcodeReplacerPost, cf.shortcodeReplacerPost)
	}
}
