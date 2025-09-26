package deps

import (
	"fmt"
	"testing"
)

func TestStartErrorCollector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &globalErrHandler{}
	ch := e.StartErrorCollector()

	if ch == nil {
		t.Errorf("StartErrorCollector() should return a non-nil channel")
	}

	if cap(ch) != 10 {
		t.Errorf("StartErrorCollector() should return a buffered channel with a capacity of 10")
	}
}
