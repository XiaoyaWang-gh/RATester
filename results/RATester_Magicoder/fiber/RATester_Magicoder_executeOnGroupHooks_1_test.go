package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnGroupHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	hooks := &Hooks{app: app}

	group := &Group{
		app:    app,
		Prefix: "/test",
	}

	hooks.OnGroup(func(g Group) error {
		if g.Prefix != "/test" {
			t.Errorf("Expected prefix /test, got %s", g.Prefix)
		}
		return nil
	})

	err := hooks.executeOnGroupHooks(*group)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
