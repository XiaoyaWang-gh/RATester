package web

import (
	"fmt"
	"testing"
)

func TestviewPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		ViewPath: "test_path",
	}

	path := ctrl.viewPath()
	if path != "test_path" {
		t.Errorf("Expected viewPath to be 'test_path', but got '%s'", path)
	}

	ctrl.ViewPath = ""
	path = ctrl.viewPath()
	if path != BConfig.WebConfig.ViewsPath {
		t.Errorf("Expected viewPath to be '%s', but got '%s'", BConfig.WebConfig.ViewsPath, path)
	}
}
