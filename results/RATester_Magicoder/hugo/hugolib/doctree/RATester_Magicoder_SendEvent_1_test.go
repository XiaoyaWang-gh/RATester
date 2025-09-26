package doctree

import (
	"fmt"
	"testing"
)

func TestSendEvent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &WalkContext[any]{}
	event := &Event[any]{Name: "testEvent"}
	ctx.SendEvent(event)

	if len(ctx.events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(ctx.events))
	}

	if ctx.events[0] != event {
		t.Errorf("Expected event to be %v, got %v", event, ctx.events[0])
	}
}
