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

	b := BuildConfig{
		List:             "always",
		Render:           "always",
		PublishResources: true,
		set:              true,
	}

	b.Disable()

	if b.List != "never" {
		t.Errorf("Expected List to be 'never', but got %s", b.List)
	}

	if b.Render != "never" {
		t.Errorf("Expected Render to be 'never', but got %s", b.Render)
	}

	if b.PublishResources != false {
		t.Errorf("Expected PublishResources to be false, but got %t", b.PublishResources)
	}

	if b.set != false {
		t.Errorf("Expected set to be false, but got %t", b.set)
	}
}
