package fiber

import (
	"fmt"
	"testing"
)

func TestMountPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			mountPath: "/test",
		},
	}

	expected := "/test"
	result := app.MountPath()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
