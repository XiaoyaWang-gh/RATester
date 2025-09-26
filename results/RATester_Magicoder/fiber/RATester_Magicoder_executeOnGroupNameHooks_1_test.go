package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnGroupNameHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	hooks := &Hooks{
		app: app,
		onGroupName: []func(Group) error{
			func(group Group) error {
				// Test logic here
				return nil
			},
		},
	}

	group := Group{
		app:    app,
		Prefix: "/test",
	}

	err := hooks.executeOnGroupNameHooks(group)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
