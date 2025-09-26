package client

import (
	"fmt"
	"testing"
)

func TestSetPathParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		path: &PathParam{},
	}

	params := map[string]string{
		"id":   "1",
		"name": "John",
	}

	r.SetPathParams(params)

	if len(*r.path) != len(params) {
		t.Errorf("Expected path params length %d, got %d", len(params), len(*r.path))
	}

	for k, v := range params {
		if (*r.path)[k] != v {
			t.Errorf("Expected path param %s to be %s, got %s", k, v, (*r.path)[k])
		}
	}
}
