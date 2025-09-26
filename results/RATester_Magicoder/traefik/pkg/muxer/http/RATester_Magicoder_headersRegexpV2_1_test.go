package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestheadersRegexpV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := headersRegexpV2(tree, "key1", "value1", "key2", "value2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("key1", "value1")
	req.Header.Set("key2", "value2")

	if !tree.match(req) {
		t.Error("Expected request to match")
	}

	req, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("key1", "wrong")
	req.Header.Set("key2", "value2")

	if tree.match(req) {
		t.Error("Expected request not to match")
	}
}
