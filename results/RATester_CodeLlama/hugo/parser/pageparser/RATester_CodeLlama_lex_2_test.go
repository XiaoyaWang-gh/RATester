package pageparser

import (
	"fmt"
	"testing"
)

func TestLex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &sectionHandlers{}
	origin := func(l *pageLexer) stateFunc {
		return nil
	}
	next := s.lex(origin)
	if next != nil {
		t.Errorf("expected nil, got %v", next)
	}
}
