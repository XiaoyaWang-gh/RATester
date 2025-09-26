package doctree

import (
	"fmt"
	"testing"
)

func TestStopPropagation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	event := &Event[int]{}
	event.StopPropagation()
	if !event.stopPropagation {
		t.Error("Expected stopPropagation to be true, but it was false")
	}
}
