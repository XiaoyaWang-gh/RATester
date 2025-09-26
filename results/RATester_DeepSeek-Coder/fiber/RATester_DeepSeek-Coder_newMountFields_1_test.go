package fiber

import (
	"fmt"
	"testing"
)

func TestNewMountFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	mountFields := newMountFields(app)

	if mountFields.appList[""] != app {
		t.Errorf("Expected appList[\"\"] to be %v, got %v", app, mountFields.appList[""])
	}

	if len(mountFields.appListKeys) != 0 {
		t.Errorf("Expected appListKeys to be empty, got %v", mountFields.appListKeys)
	}
}
