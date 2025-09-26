package fiber

import (
	"fmt"
	"testing"
)

func TeststartupProcess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}

	app.startupProcess()

	// Assertions
	if app.mountFields == nil {
		t.Errorf("Expected mountFields to be set, but it was nil")
	}
	if app.stack == nil {
		t.Errorf("Expected stack to be set, but it was nil")
	}
	if app.treeStack == nil {
		t.Errorf("Expected treeStack to be set, but it was nil")
	}
	if app.routesCount != 0 {
		t.Errorf("Expected routesCount to be 0, but it was %d", app.routesCount)
	}
	if app.handlersCount != 0 {
		t.Errorf("Expected handlersCount to be 0, but it was %d", app.handlersCount)
	}
	if !app.routesRefreshed {
		t.Errorf("Expected routesRefreshed to be true, but it was false")
	}
}
