package pagemeta

import (
	"fmt"
	"testing"
)

func TestDisable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildConfig{}
	b.Disable()
	if b.List != Never {
		t.Errorf("b.List = %q, want %q", b.List, Never)
	}
	if b.Render != Never {
		t.Errorf("b.Render = %q, want %q", b.Render, Never)
	}
	if b.PublishResources {
		t.Errorf("b.PublishResources = %v, want %v", b.PublishResources, false)
	}
	if !b.set {
		t.Errorf("b.set = %v, want %v", b.set, true)
	}
}
