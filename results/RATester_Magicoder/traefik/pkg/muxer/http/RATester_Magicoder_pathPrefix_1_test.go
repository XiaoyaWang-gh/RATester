package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestpathPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}

	err := pathPrefix(tree, "/test")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !tree.match(req) {
		t.Errorf("Expected match, but it didn't")
	}

	req, err = http.NewRequest("GET", "/nottest", nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if tree.match(req) {
		t.Errorf("Expected no match, but it did")
	}
}
