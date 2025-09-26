package fiber

import (
	"errors"
	"fmt"
	"testing"
)

func TestexecuteOnShutdownHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hooks := &Hooks{
		onShutdown: []func() error{
			func() error {
				return nil
			},
			func() error {
				return errors.New("test error")
			},
		},
	}

	hooks.executeOnShutdownHooks()

	// Test case for error
	hooks.onShutdown[1] = func() error {
		return errors.New("test error")
	}

	hooks.executeOnShutdownHooks()
}
