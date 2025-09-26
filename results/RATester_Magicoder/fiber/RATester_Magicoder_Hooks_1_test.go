package fiber

import (
	"fmt"
	"testing"
)

func TestHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	hooks := &Hooks{}
	app.hooks = hooks

	result := app.Hooks()

	if result != hooks {
		t.Errorf("Expected %v, got %v", hooks, result)
	}
}
