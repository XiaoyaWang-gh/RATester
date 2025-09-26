package httpcache

import (
	"fmt"
	"testing"
)

func TestCompilePredicate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gm := &GlobMatcher{
		Excludes: []string{"/foo/bar", "/foo/baz"},
		Includes: []string{"/foo/bar", "/foo/baz"},
	}

	p, err := gm.CompilePredicate()
	if err != nil {
		t.Fatal(err)
	}

	if !p("/foo/bar") {
		t.Errorf("expected %q to match", "/foo/bar")
	}

	if !p("/foo/baz") {
		t.Errorf("expected %q to match", "/foo/baz")
	}

	if p("/foo/qux") {
		t.Errorf("expected %q to not match", "/foo/qux")
	}
}
