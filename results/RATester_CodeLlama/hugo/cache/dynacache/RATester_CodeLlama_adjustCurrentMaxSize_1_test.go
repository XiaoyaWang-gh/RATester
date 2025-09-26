package dynacache

import (
	"fmt"
	"testing"
)

func TestAdjustCurrentMaxSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &stats{}
	s.opts.MaxSize = 100
	s.opts.MinMaxSize = 10
	s.adjustmentFactor = 0.5
	s.currentMaxSize = 100
	changed := s.adjustCurrentMaxSize()
	if changed != true {
		t.Errorf("adjustCurrentMaxSize() = %v, want %v", changed, true)
	}
	if s.currentMaxSize != 50 {
		t.Errorf("adjustCurrentMaxSize() = %v, want %v", s.currentMaxSize, 50)
	}
}
