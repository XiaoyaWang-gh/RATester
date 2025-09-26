package identity

import (
	"fmt"
	"testing"
)

func TestNewFinder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := FinderConfig{Exact: true}
	finder := NewFinder(cfg)

	if finder.cfg.Exact != cfg.Exact {
		t.Errorf("Expected Exact to be %v, got %v", cfg.Exact, finder.cfg.Exact)
	}

	if finder.answers == nil {
		t.Errorf("Expected answers map to be initialized")
	}

	if finder.seenFindOnce == nil {
		t.Errorf("Expected seenFindOnce map to be initialized")
	}
}
