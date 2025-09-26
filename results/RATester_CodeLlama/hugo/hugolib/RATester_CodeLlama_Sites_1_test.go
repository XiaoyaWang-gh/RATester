package hugolib

import (
	"fmt"
	"testing"
)

func TestSites_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	s := &Site{}
	sites := s.Sites()
	if len(sites) != 0 {
		t.Errorf("Expected empty sites, got %d", len(sites))
	}
}
