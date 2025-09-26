package httplib

import (
	"fmt"
	"testing"
)

func TestParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		url: "http://example.com",
		params: map[string][]string{
			"key1": {"value1"},
		},
	}

	req.Param("key2", "value2")

	if len(req.params) != 2 {
		t.Errorf("Expected 2 params, got %d", len(req.params))
	}

	if req.params["key2"][0] != "value2" {
		t.Errorf("Expected value2 for key2, got %s", req.params["key2"][0])
	}
}
