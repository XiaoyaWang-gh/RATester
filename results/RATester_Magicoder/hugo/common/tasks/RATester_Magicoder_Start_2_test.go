package tasks

import (
	"fmt"
	"testing"
)

func TestStart_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		RunImmediately: true,
	}

	err := r.Start()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !r.started {
		t.Errorf("Expected started to be true, got false")
	}

	if r.quit == nil {
		t.Errorf("Expected quit to be initialized, got nil")
	}
}
