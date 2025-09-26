package echo

import (
	"fmt"
	"testing"
)

func TestFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	file := "file"
	get := func(string, HandlerFunc, ...MiddlewareFunc) *Route {
		return &Route{}
	}
	m := []MiddlewareFunc{}
	r := common{}.file(path, file, get, m...)
	if r.Method != "GET" {
		t.Errorf("Method = %v, want %v", r.Method, "GET")
	}
	if r.Path != path {
		t.Errorf("Path = %v, want %v", r.Path, path)
	}
	if r.Name != "" {
		t.Errorf("Name = %v, want %v", r.Name, "")
	}
}
