package web

import (
	"fmt"
	"testing"
)

func TestSetViewsPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "test"
	app := SetViewsPath(path)
	if app.Cfg.WebConfig.ViewsPath != path {
		t.Errorf("SetViewsPath failed")
	}
}
