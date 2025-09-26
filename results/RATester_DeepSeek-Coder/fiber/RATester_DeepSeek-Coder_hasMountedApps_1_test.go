package fiber

import (
	"fmt"
	"testing"
)

func TestHasMountedApps_1(t *testing.T) {
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

	hasMountedApps := app.hasMountedApps()
	if !hasMountedApps {
		t.Errorf("Expected true, got false")
	}

	app = &App{
		mountFields: &mountFields{
			appList: map[string]*App{},
		},
	}

	hasMountedApps = app.hasMountedApps()
	if hasMountedApps {
		t.Errorf("Expected false, got true")
	}
}
