package fiber

import (
	"fmt"
	"testing"
)

func TesthasMountedApps_1(t *testing.T) {
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
			},
		},
	}

	result := app.hasMountedApps()

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}
