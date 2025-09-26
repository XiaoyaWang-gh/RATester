package hugo

import (
	"fmt"
	"testing"
)

func TestDeps_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := HugoInfo{
		deps: []*Dependency{
			{Path: "dep1", Version: "1.0.0"},
			{Path: "dep2", Version: "2.0.0"},
		},
	}

	deps := info.Deps()

	if len(deps) != 2 {
		t.Errorf("Expected 2 dependencies, got %d", len(deps))
	}

	if deps[0].Path != "dep1" || deps[0].Version != "1.0.0" {
		t.Errorf("Expected dep1@1.0.0, got %s@%s", deps[0].Path, deps[0].Version)
	}

	if deps[1].Path != "dep2" || deps[1].Version != "2.0.0" {
		t.Errorf("Expected dep2@2.0.0, got %s@%s", deps[1].Path, deps[1].Version)
	}
}
