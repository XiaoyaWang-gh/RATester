package pageparser

import (
	"fmt"
	"testing"
)

func TestSkip_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &sectionHandler{}
	s.skipFunc = func(l *pageLexer) int {
		return -1
	}
	if s.skip() != -1 {
		t.Errorf("Expected -1, got %d", s.skip())
	}
}
