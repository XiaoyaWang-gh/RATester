package echo

import (
	"fmt"
	"testing"
)

func TestFindRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	r := &Router{}
	e.routers = make(map[string]*Router)
	e.routers["test"] = r

	result := e.findRouter("test")

	if result != r {
		t.Errorf("Expected %v, got %v", r, result)
	}
}
