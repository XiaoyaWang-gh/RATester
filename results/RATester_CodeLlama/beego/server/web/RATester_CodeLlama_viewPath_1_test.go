package web

import (
	"fmt"
	"testing"
)

func TestViewPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.ViewPath = "viewPath"
	if c.viewPath() != "viewPath" {
		t.Errorf("viewPath() = %v, want %v", c.viewPath(), "viewPath")
	}
}
