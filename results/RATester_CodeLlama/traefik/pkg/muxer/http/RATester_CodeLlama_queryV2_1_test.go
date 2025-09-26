package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQueryV2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	err := queryV2(tree, "query1=value1", "query2=value2")
	if err != nil {
		t.Errorf("queryV2() error = %v", err)
		return
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("http.NewRequest() error = %v", err)
		return
	}

	req.URL.RawQuery = "query1=value1&query2=value2"
	if !tree.match(req) {
		t.Errorf("tree.match() = %v, want %v", false, true)
	}
}
