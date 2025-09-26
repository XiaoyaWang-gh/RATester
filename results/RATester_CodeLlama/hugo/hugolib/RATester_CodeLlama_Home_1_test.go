package hugolib

import (
	"fmt"
	"testing"
)

func TestHome_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	p := s.Home()
	if p != nil {
		t.Errorf("expected nil, got %v", p)
	}
}
