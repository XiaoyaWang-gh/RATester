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

	t.Parallel()
	cfg := FinderConfig{Exact: true}
	f := NewFinder(cfg)
	if f.cfg.Exact != true {
		t.Errorf("Expected Exact to be true, got %v", f.cfg.Exact)
	}
}
