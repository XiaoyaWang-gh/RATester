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
	errorsChan := e.StartErrorCollector()

	if errorsChan == nil {
		t.Error("Expected a non-nil error channel")
	}

	if cap(errorsChan) != 10 {
		t.Error("Expected a buffer size of 10")
	}
}
