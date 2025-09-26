package deps

import (
	"fmt"
	"testing"
)

func TestStopErrorCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &globalErrHandler{}
	e.buildErrors = make(chan error)
	e.quit = make(chan struct{})
	e.StopErrorCollector()
	if e.buildErrors != nil {
		t.Errorf("buildErrors should be nil")
	}
	if e.quit != nil {
		t.Errorf("quit should be nil")
	}
}
