package fiber

import (
	"fmt"
	"sync"
	"testing"
)

func TestStartupProcess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := &App{
		mutex: sync.Mutex{},
	}

	result := app.startupProcess()

	if result != app {
		t.Errorf("Expected %v, got %v", app, result)
	}
}
