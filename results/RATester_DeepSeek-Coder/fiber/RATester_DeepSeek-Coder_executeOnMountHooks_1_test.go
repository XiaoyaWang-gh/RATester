package fiber

import (
	"errors"
	"fmt"
	"testing"
)

func TestExecuteOnMountHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		hooks: &Hooks{
			onMount: []func(*App) error{
				func(app *App) error {
					return nil
				},
				func(app *App) error {
					return errors.New("test error")
				},
			},
		},
	}

	err := app.hooks.executeOnMountHooks(app)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	app.hooks.onMount = []func(*App) error{
		func(app *App) error {
			return errors.New("test error")
		},
	}

	err = app.hooks.executeOnMountHooks(app)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
