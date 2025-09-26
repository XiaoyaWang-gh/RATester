package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnMountHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		hooks: &Hooks{
			onMount: []func(*App) error{
				func(app *App) error {
					// do something
					return nil
				},
			},
		},
	}

	err := app.hooks.executeOnMountHooks(app)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
