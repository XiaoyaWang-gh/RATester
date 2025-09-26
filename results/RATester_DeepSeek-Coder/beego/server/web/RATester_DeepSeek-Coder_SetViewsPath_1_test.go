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

	testPath := "/test/path"
	SetViewsPath(testPath)
	if BConfig.WebConfig.ViewsPath != testPath {
		t.Errorf("Expected ViewsPath to be %s, got %s", testPath, BConfig.WebConfig.ViewsPath)
	}
}
