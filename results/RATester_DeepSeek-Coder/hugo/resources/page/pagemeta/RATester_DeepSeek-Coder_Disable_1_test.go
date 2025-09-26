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

	b := &BuildConfig{
		List:             "always",
		Render:           "always",
		PublishResources: true,
		set:              true,
	}

	b.Disable()

	if b.List != Never {
		t.Errorf("Expected List to be %s, got %s", Never, b.List)
	}

	if b.Render != Never {
		t.Errorf("Expected Render to be %s, got %s", Never, b.Render)
	}

	if b.PublishResources != false {
		t.Errorf("Expected PublishResources to be false, got %t", b.PublishResources)
	}

	if b.set != true {
		t.Errorf("Expected set to be true, got %t", b.set)
	}
}
