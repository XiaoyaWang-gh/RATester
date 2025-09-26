package helpers

import (
	"fmt"
	"testing"
)

func TestMakePathsSanitized_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	paths := []string{"/foo/bar", "/foo/bar/"}
	p.MakePathsSanitized(paths)
	if paths[0] != "/foo/bar" {
		t.Errorf("Expected %q but got %q", "/foo/bar", paths[0])
	}
	if paths[1] != "/foo/bar" {
		t.Errorf("Expected %q but got %q", "/foo/bar", paths[1])
	}
}
