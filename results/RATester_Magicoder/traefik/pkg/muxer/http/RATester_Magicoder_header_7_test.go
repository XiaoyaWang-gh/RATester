package http

import (
	"fmt"
	"net/http"
	"testing"
)

func Testheader_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	headers := []string{"Content-Type", "application/json"}

	err := header(tree, headers...)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	if !tree.match(req) {
		t.Error("Expected match, but got no match")
	}

	req.Header.Set("Content-Type", "text/plain")

	if tree.match(req) {
		t.Error("Expected no match, but got match")
	}
}
