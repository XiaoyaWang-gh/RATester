package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQueryRegexp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	queries := []string{"key", "value"}
	err := queryRegexp(tree, queries...)
	if err != nil {
		t.Errorf("queryRegexp() error = %v, wantErr nil", err)
		return
	}

	req, err := http.NewRequest("GET", "http://example.com?key=value", nil)
	if err != nil {
		t.Errorf("http.NewRequest() error = %v, wantErr nil", err)
		return
	}

	if !tree.match(req) {
		t.Errorf("tree.match() = %v, want %v", false, true)
	}
}
