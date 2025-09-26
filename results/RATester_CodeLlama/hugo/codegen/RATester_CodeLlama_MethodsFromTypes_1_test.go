package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMethodsFromTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Inspector{
		ProjectRootDir: "github.com/gohugoio/hugo",
	}

	var include []reflect.Type
	var exclude []reflect.Type

	methods := c.MethodsFromTypes(include, exclude)

	if len(methods) != 0 {
		t.Errorf("Expected 0 methods, got %d", len(methods))
	}
}
