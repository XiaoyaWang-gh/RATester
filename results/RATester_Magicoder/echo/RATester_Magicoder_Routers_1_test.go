package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	expectedRouters := map[string]*Router{
		"router1": &Router{},
		"router2": &Router{},
	}
	e.routers = expectedRouters

	result := e.Routers()

	if !reflect.DeepEqual(result, expectedRouters) {
		t.Errorf("Expected %v, but got %v", expectedRouters, result)
	}
}
