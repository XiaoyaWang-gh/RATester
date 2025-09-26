package fiber

import (
	"fmt"
	"testing"
)

func TestnewMountFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	mountFields := newMountFields(app)

	if mountFields.appList == nil {
		t.Error("Expected appList to be initialized")
	}

	if mountFields.appListKeys == nil {
		t.Error("Expected appListKeys to be initialized")
	}

	if mountFields.appList[""] != app {
		t.Error("Expected app to be added to appList")
	}
}
