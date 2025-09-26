package hugolib

import (
	"fmt"
	"testing"
	"time"
)

func TestLastChange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.lastmod = time.Now()
	if s.LastChange() != s.lastmod {
		t.Errorf("LastChange() = %v, want %v", s.LastChange(), s.lastmod)
	}
}
