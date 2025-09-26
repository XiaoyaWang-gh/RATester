package fiber

import (
	"fmt"
	"testing"
)

func TestAppendSubAppLists_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			appList: make(map[string]*App),
		},
	}

	subApp1 := &App{
		mountFields: &mountFields{
			appList: make(map[string]*App),
		},
	}

	subApp2 := &App{
		mountFields: &mountFields{
			appList: make(map[string]*App),
		},
	}

	appList := map[string]*App{
		"sub1": subApp1,
		"sub2": subApp2,
	}

	app.appendSubAppLists(appList)

	if len(app.mountFields.appList) != 2 {
		t.Errorf("Expected 2 sub apps, got %d", len(app.mountFields.appList))
	}

	if _, ok := app.mountFields.appList["sub1"]; !ok {
		t.Errorf("Expected subApp1 to be in app.mountFields.appList")
	}

	if _, ok := app.mountFields.appList["sub2"]; !ok {
		t.Errorf("Expected subApp2 to be in app.mountFields.appList")
	}
}
