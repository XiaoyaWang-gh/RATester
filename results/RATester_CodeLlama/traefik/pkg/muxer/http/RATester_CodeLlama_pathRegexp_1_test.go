package http

import (
	"fmt"
	"testing"
)

func TestPathRegexp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	paths := []string{"/path/to/resource"}
	err := pathRegexp(tree, paths...)
	if err != nil {
		t.Errorf("pathRegexp() error = %v, want nil", err)
		return
	}
	if tree.matcher == nil {
		t.Errorf("pathRegexp() matcher = nil, want not nil")
	}
}
