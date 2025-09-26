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
		funcs: make(map[string]*Func),
	}

	if err := r.Start(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if r.started {
		t.Error("expected started to be false")
	}

	if r.quit != nil {
		t.Error("expected quit to be nil")
	}

	r.started = true
	r.quit = make(chan struct{})

	if err := r.Start(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !r.started {
		t.Error("expected started to be true")
	}

	if r.quit == nil {
		t.Error("expected quit to be non-nil")
	}
}
