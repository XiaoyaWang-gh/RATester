package gin

import (
	"fmt"
	"testing"
)

func TestHandleContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{}
	context := &Context{}

	oldIndexValue := context.index
	context.reset()
	engine.handleHTTPRequest(context)

	if context.index != oldIndexValue {
		t.Errorf("Expected index to be %d, but got %d", oldIndexValue, context.index)
	}
}
