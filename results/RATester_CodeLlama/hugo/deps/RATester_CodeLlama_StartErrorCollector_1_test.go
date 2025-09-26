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
	e.StartErrorCollector()
	if e.buildErrors == nil {
		t.Error("buildErrors should not be nil")
	}
	if e.quit == nil {
		t.Error("quit should not be nil")
	}
}
