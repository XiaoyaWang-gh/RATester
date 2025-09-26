package fiber

import (
	"fmt"
	"testing"
)

func TestHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	handler := app.Handler()

	if handler == nil {
		t.Error("Handler is nil")
	}
}
