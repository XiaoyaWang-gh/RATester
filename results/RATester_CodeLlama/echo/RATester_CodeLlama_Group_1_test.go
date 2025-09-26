package echo

import (
	"fmt"
	"testing"
)

func TestGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &Group{}
	prefix := "prefix"
	middleware := []MiddlewareFunc{}
	sg := g.Group(prefix, middleware...)
	if sg == nil {
		t.Errorf("sg is nil")
	}
}
