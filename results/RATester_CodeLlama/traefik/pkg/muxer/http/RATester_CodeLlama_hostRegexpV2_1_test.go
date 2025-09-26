package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHostRegexpV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	hosts := []string{"example.com", "example.net"}
	err := hostRegexpV2(tree, hosts...)
	if err != nil {
		t.Errorf("hostRegexpV2() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("hostRegexpV2() matcher = nil")
		return
	}

	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Errorf("hostRegexpV2() http.NewRequest() error = %v", err)
		return
	}

	if !tree.match(req) {
		t.Errorf("hostRegexpV2() match() = false, want true")
	}

	req, err = http.NewRequest("GET", "http://example.net/", nil)
	if err != nil {
		t.Errorf("hostRegexpV2() http.NewRequest() error = %v", err)
		return
	}

	if !tree.match(req) {
		t.Errorf("hostRegexpV2() match() = false, want true")
	}

	req, err = http.NewRequest("GET", "http://example.org/", nil)
	if err != nil {
		t.Errorf("hostRegexpV2() http.NewRequest() error = %v", err)
		return
	}

	if tree.match(req) {
		t.Errorf("hostRegexpV2() match() = true, want false")
	}
}
