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

	app := New()
	hooks := app.Hooks()

	if hooks == nil {
		t.Errorf("Expected Hooks to not be nil")
	}
}
