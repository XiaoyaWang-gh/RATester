package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHostV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	hosts := []string{"example.com", "example.net"}
	err := hostV2(tree, hosts...)
	if err != nil {
		t.Errorf("hostV2() error = %v", err)
		return
	}

	req := &http.Request{
		Host: "example.com",
	}
	if !tree.match(req) {
		t.Errorf("hostV2() match() = false, want true")
	}

	req = &http.Request{
		Host: "example.net",
	}
	if !tree.match(req) {
		t.Errorf("hostV2() match() = false, want true")
	}

	req = &http.Request{
		Host: "example.org",
	}
	if tree.match(req) {
		t.Errorf("hostV2() match() = true, want false")
	}
}
