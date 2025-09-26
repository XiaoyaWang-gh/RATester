package fiber

import (
	"fmt"
	"testing"
)

func TestnewHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	hooks := newHooks(app)

	if hooks.app != app {
		t.Errorf("Expected app to be %v, but got %v", app, hooks.app)
	}

	if len(hooks.onRoute) != 0 {
		t.Errorf("Expected onRoute to be empty, but got %v", hooks.onRoute)
	}

	if len(hooks.onName) != 0 {
		t.Errorf("Expected onName to be empty, but got %v", hooks.onName)
	}

	if len(hooks.onGroup) != 0 {
		t.Errorf("Expected onGroup to be empty, but got %v", hooks.onGroup)
	}

	if len(hooks.onGroupName) != 0 {
		t.Errorf("Expected onGroupName to be empty, but got %v", hooks.onGroupName)
	}

	if len(hooks.onListen) != 0 {
		t.Errorf("Expected onListen to be empty, but got %v", hooks.onListen)
	}

	if len(hooks.onShutdown) != 0 {
		t.Errorf("Expected onShutdown to be empty, but got %v", hooks.onShutdown)
	}

	if len(hooks.onFork) != 0 {
		t.Errorf("Expected onFork to be empty, but got %v", hooks.onFork)
	}

	if len(hooks.onMount) != 0 {
		t.Errorf("Expected onMount to be empty, but got %v", hooks.onMount)
	}
}
