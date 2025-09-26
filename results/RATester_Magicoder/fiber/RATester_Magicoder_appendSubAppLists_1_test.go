package fiber

import (
	"fmt"
	"testing"
)

func TestappendSubAppLists_1(t *testing.T) {
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

	subApp := &App{
		mountFields: &mountFields{
			appList: map[string]*App{
				"sub": {},
			},
		},
	}

	appList := map[string]*App{
		"": subApp,
	}

	app.appendSubAppLists(appList)

	if _, ok := app.mountFields.appList["sub"]; !ok {
		t.Error("Expected subApp to be added to app.mountFields.appList")
	}
}
