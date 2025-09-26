package maps

import (
	"fmt"
	"testing"
)

func TestScratch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := scratcher{s: &Scratch{}}
	if s.Scratch() != s.s {
		t.Errorf("Scratch() = %v, want %v", s.Scratch(), s.s)
	}
}
