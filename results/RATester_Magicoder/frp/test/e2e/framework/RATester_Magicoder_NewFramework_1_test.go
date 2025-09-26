package framework

import (
	"fmt"
	"testing"
)

func TestNewFramework_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	opt := Options{
		TotalParallelNode: 10,
		CurrentNodeIndex:  5,
		FromPortIndex:     1000,
		ToPortIndex:       2000,
	}

	f := NewFramework(opt)

	if f == nil {
		t.Errorf("NewFramework() should not return nil")
	}

	if f.portAllocator == nil {
		t.Errorf("portAllocator should not be nil")
	}

	if len(f.usedPorts) != 0 {
		t.Errorf("usedPorts should be empty")
	}

	if f.beforeEachStarted {
		t.Errorf("beforeEachStarted should be false")
	}
}
