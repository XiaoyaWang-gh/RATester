package hugolib

import (
	"fmt"
	"testing"
)

func TestprepareInits_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{
		init: &siteInit{},
	}

	s.prepareInits()

	if s.init.prevNext == nil {
		t.Error("prevNext is nil")
	}

	if s.init.prevNextInSection == nil {
		t.Error("prevNextInSection is nil")
	}

	if s.init.menus == nil {
		t.Error("menus is nil")
	}

	if s.init.taxonomies == nil {
		t.Error("taxonomies is nil")
	}
}
