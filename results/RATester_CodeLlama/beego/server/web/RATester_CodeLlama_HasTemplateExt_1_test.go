package web

import (
	"fmt"
	"testing"
)

func TestHasTemplateExt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var paths string
	if HasTemplateExt(paths) {
		t.Errorf("HasTemplateExt(%q) = true, want false", paths)
	}
}
