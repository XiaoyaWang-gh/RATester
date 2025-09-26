package http

import (
	"fmt"
	"net/http"
	"testing"
)

func Testmethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := method(tree, "GET")
	if err != nil {
		t.Errorf("method() error = %v", err)
		return
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("http.NewRequest() error = %v", err)
		return
	}

	if !tree.match(req) {
		t.Errorf("tree.match() = false, want true")
	}
}
