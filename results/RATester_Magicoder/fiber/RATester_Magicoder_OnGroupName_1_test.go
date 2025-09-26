package fiber

import (
	"fmt"
	"testing"
)

func TestOnGroupName_1(t *testing.T) {
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

	hooks.OnGroupName(handler)

	if len(hooks.onGroupName) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(hooks.onGroupName))
	}

	hooks.executeOnGroupNameHooks(Group{})

	// Test logic here
}
