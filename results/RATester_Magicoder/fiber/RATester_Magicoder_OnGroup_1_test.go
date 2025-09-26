package fiber

import (
	"fmt"
	"testing"
)

func TestOnGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handler := func(group Group) error {
		// Test logic here
		return nil
	}

	hooks.OnGroup(handler)

	if len(hooks.onGroup) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(hooks.onGroup))
	}

	hooks.executeOnGroupHooks(Group{})
}
