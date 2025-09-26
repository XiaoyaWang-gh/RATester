package http

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	paths := []string{"/path"}
	err := path(tree, paths...)
	if err != nil {
		t.Errorf("path() error = %v", err)
		return
	}
	if tree.matcher == nil {
		t.Errorf("path() matcher is nil")
		return
	}
	req := &http.Request{
		URL: &url.URL{
			Path: "/path",
		},
	}
	if !tree.match(req) {
		t.Errorf("path() matcher does not match")
	}
}
