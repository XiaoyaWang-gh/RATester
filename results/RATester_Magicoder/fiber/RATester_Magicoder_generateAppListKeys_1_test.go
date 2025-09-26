package fiber

import (
	"fmt"
	"testing"
)

func TestgenerateAppListKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			appList: map[string]*App{
				"app1": {},
				"app2": {},
				"app3": {},
			},
		},
	}

	app.generateAppListKeys()

	if len(app.mountFields.appListKeys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(app.mountFields.appListKeys))
	}

	if app.mountFields.appListKeys[0] != "app1" {
		t.Errorf("Expected 'app1', got %s", app.mountFields.appListKeys[0])
	}

	if app.mountFields.appListKeys[1] != "app2" {
		t.Errorf("Expected 'app2', got %s", app.mountFields.appListKeys[1])
	}

	if app.mountFields.appListKeys[2] != "app3" {
		t.Errorf("Expected 'app3', got %s", app.mountFields.appListKeys[2])
	}
}
