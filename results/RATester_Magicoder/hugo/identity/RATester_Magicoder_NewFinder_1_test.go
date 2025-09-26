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

	if finder == nil {
		t.Error("NewFinder returned nil")
	}

	if finder.cfg.Exact != cfg.Exact {
		t.Errorf("NewFinder did not set the correct config. Expected: %v, Got: %v", cfg.Exact, finder.cfg.Exact)
	}
}
