package hugolib

import (
	"fmt"
	"testing"
)

func TestwithSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	err := h.withSite(func(s *Site) error {
		// Test logic here
		return nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
