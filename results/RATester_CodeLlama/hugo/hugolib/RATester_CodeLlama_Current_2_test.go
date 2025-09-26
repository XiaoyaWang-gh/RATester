package hugolib

import (
	"fmt"
	"testing"
)

func TestCurrent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.h = &HugoSites{}
	s.h.currentSite = s
	if s.Current() != s {
		t.Errorf("expected %p, got %p", s, s.Current())
	}
}
