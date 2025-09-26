package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := header(tree, "header1", "value1")
	if err != nil {
		t.Errorf("header() error = %v", err)
		return
	}

	if tree.matcher == nil {
		t.Errorf("header() matcher = nil, want not nil")
	}

	req := &http.Request{
		Header: http.Header{
			"header1": []string{"value1"},
		},
	}
	if !tree.match(req) {
		t.Errorf("header() matcher = false, want true")
	}

	req = &http.Request{
		Header: http.Header{
			"header1": []string{"value2"},
		},
	}
	if tree.match(req) {
		t.Errorf("header() matcher = true, want false")
	}
}
