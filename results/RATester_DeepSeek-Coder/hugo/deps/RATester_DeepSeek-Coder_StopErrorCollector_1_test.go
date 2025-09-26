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

	e := &globalErrHandler{
		buildErrors: make(chan error),
		quit:        make(chan struct{}),
	}

	e.StopErrorCollector()

	select {
	case <-e.quit:
		t.Error("quit channel should not be closed")
	default:
	}

	select {
	case <-e.buildErrors:
		t.Error("buildErrors channel should not be closed")
	default:
	}
}
