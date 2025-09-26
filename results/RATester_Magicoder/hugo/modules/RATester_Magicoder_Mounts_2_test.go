package modules

import (
	"fmt"
	"testing"
)

func TestMounts_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	adapter := &moduleAdapter{
		mounts: []Mount{
			{Source: "source1", Target: "target1"},
			{Source: "source2", Target: "target2"},
		},
	}

	mounts := adapter.Mounts()

	if len(mounts) != 2 {
		t.Errorf("Expected 2 mounts, got %d", len(mounts))
	}

	if mounts[0].Source != "source1" || mounts[0].Target != "target1" {
		t.Errorf("Expected first mount to be source1 and target1, got %s and %s", mounts[0].Source, mounts[0].Target)
	}

	if mounts[1].Source != "source2" || mounts[1].Target != "target2" {
		t.Errorf("Expected second mount to be source2 and target2, got %s and %s", mounts[1].Source, mounts[1].Target)
	}
}
