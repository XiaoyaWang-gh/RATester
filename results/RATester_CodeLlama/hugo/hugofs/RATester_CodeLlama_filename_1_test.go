package hugofs

import (
	"fmt"
	"testing"
)

func TestFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := RootMapping{
		From: "blog",
		To:   "/Users/user/go/src/github.com/gohugoio/hugo/exampleSite/content/blog",
	}
	name := "blog/my-post.md"
	expected := "/Users/user/go/src/github.com/gohugoio/hugo/exampleSite/content/blog/my-post.md"
	if r.filename(name) != expected {
		t.Errorf("expected %q but got %q", expected, r.filename(name))
	}
}
